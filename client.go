package mqttmon

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client is a middleman between the mqtt client connection and the hub.
type Client struct {
	// Broker Information
	broker *Broker

	//  Logger Entry
	logEntry *logrus.Entry

	// Publish payload
	Publish chan []byte

	// Subscribe to broker
	Subscribe chan []byte

	// Unsubscribe to broker
	Unsubscribe chan []byte

	// Close channel
	Close chan bool

	// listener
	// listeners map[string][]chan []byte
	listeners map[string]chan []byte
}

// NewClient create new client
func NewClient(
	broker *Broker,
	logEntry *logrus.Entry,
) *Client {
	return &Client{
		broker:      broker,
		logEntry:    logEntry,
		Publish:     make(chan []byte),
		Subscribe:   make(chan []byte),
		Unsubscribe: make(chan []byte),
		Close:       make(chan bool),
	}
}

// AddListener add event listener
// func (c *Client) AddListener(e string, ch chan []byte) {
// 	if c.listeners == nil {
// 		c.listeners = make(map[string][]chan []byte)
// 	}
// 	c.listeners[e] = append(c.listeners[e], ch)
// }
func (c *Client) AddListener(e string, ch chan []byte) {
	if c.listeners == nil {
		c.listeners = make(map[string]chan []byte)
	}
	c.listeners[e] = ch
}

// RemoveListener removes an event listener
// func (c *Client) RemoveListener(e string, ch chan []byte) {
// 	if _, ok := c.listeners[e]; ok {
// 		for i := range c.listeners[e] {
// 			if c.listeners[e][i] == ch {
// 				c.listeners[e] = append(c.listeners[e][:i], c.listeners[e][i+1:]...)
// 				break
// 			}
// 		}
// 	}
// }
func (c *Client) RemoveListener(e string) {
	delete(c.listeners, e)
}

// Emit emits an event on the Dog struct instance
// func (c *Client) emit(e string, response []byte) {
// 	if _, ok := c.listeners[e]; ok {
// 		for _, handler := range c.listeners[e] {
// 			go func(handler chan []byte) {
// 				handler <- response
// 			}(handler)
// 		}
// 	}
// }

func (c *Client) emit(e string, response []byte) {
	if listener, ok := c.listeners[e]; ok {
		fmt.Printf("Emit to : %s", e)
		go func(listener chan []byte) {
			listener <- response
		}(listener)
	}
}

// Connect to  Broker
// func (c *Client) Connect(wg *sync.WaitGroup) {
func (c *Client) Connect() {
	hostName, _ := os.Hostname()

	var clientid = ""
	if c.broker.ClientID == "nil" || c.broker.ClientID == "" {
		clientid = hostName + strconv.Itoa(time.Now().Second())
	}

	options := mqtt.NewClientOptions().AddBroker(c.broker.BrokerAddress).SetClientID(clientid).SetCleanSession(true)
	if c.broker.Username != "" && c.broker.ClientID != "nil" {
		options.SetUsername(c.broker.Username)
		if c.broker.Userpass != "" && c.broker.Userpass != "nil" {
			options.SetPassword(c.broker.Userpass)
		}
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	options.SetTLSConfig(tlsConfig)
	options.SetAutoReconnect(c.broker.AutoReconnect == 1)
	options.SetKeepAlive(time.Duration(c.broker.KeepAlive) * time.Second)
	options.SetPingTimeout(time.Duration(c.broker.PingTimeout) * time.Second)
	options.SetConnectRetryInterval(time.Duration(c.broker.ConnRetryDelay) * time.Second)

	// Connecting to Broker
	fmt.Printf("Connecting to %s\n", c.broker.BrokerAddress)
	client := mqtt.NewClient(options)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		c.logEntry.Error(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", c.broker.BrokerAddress)
	}

	// wg.Done()

	defer client.Disconnect(250)

	c.run(client)

}

func (c *Client) run(client mqtt.Client) {
	for {
		select {
		case data, ok := <-c.Publish:
			if !ok {
				fmt.Printf("Channel is closed!\n")
				return
			}

			var publication Publication

			err := json.Unmarshal(data, &publication)
			if err != nil {
				c.logEntry.Error(err)
				continue
			}

			messageJSON, err := json.Marshal(publication.Payload)
			if err != nil {
				c.logEntry.Error(err)
			}

			// fmt.Println(publish)

			if token := client.Publish(publication.Topic, byte(publication.QOS), publication.IsRetain == 1, messageJSON); token.Wait() && token.Error() != nil {
				c.logEntry.Error(token.Error())
			}
		case data, ok := <-c.Subscribe:
			if !ok {
				fmt.Printf("Channel is closed!\n")
				return
			}

			var subscription Subscription

			err := json.Unmarshal(data, &subscription)
			if err != nil {
				c.logEntry.Error(err)
				continue
			}

			if token := client.Subscribe(subscription.Topic, byte(subscription.QOS), func(client mqtt.Client, message mqtt.Message) {
				go func(message mqtt.Message) {
					// c.logEntry.Info(message)
					fmt.Printf("still emit!\n")
					c.emit(subscription.SessionID, message.Payload())
				}(message)
			}); token.Wait() && token.Error() != nil {
				c.logEntry.Error(token.Error())
			}

		case data, ok := <-c.Unsubscribe:
			if !ok {
				fmt.Printf("Channel is closed!\n")
				return
			}

			var subscription Subscription

			err := json.Unmarshal(data, &subscription)
			if err != nil {
				c.logEntry.Error(err)
				continue
			}

			if token := client.Unsubscribe(subscription.Topic); token.Wait() && token.Error() != nil {
				c.logEntry.Error(token.Error())
				// fmt.Printf("Removing listener: %s\n", subscription.SessionID)
				// timer1 := time.NewTimer(30 * time.Second)

				// go func() {
				// 	<-timer1.C
				// 	c.RemoveListener(subscription.SessionID)
				// 	fmt.Printf("Removed listener: %s\n", subscription.SessionID)
				// }()
			}
		default:
			continue
		}
	}
}
