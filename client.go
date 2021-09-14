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

	// Close channel
	Close chan bool

	// listener
	listeners map[string][]chan []byte
}

// NewClient create new client
func NewClient(
	broker *Broker,
	logEntry *logrus.Entry,
) *Client {
	return &Client{
		broker:    broker,
		logEntry:  logEntry,
		Publish:   make(chan []byte),
		Subscribe: make(chan []byte),
		Close:     make(chan bool),
	}
}

// AddListener add event listener
func (c *Client) AddListener(e string, ch chan []byte) {
	if c.listeners == nil {
		c.listeners = make(map[string][]chan []byte)
	}
	c.listeners[e] = append(c.listeners[e], ch)
}

// RemoveListener removes an event listener
func (c *Client) RemoveListener(e string, ch chan []byte) {
	if _, ok := c.listeners[e]; ok {
		for i := range c.listeners[e] {
			if c.listeners[e][i] == ch {
				c.listeners[e] = append(c.listeners[e][:i], c.listeners[e][i+1:]...)
				break
			}
		}
	}
}

// Emit emits an event on the Dog struct instance
func (c *Client) emit(e string, response []byte) {
	if _, ok := c.listeners[e]; ok {
		for _, handler := range c.listeners[e] {
			go func(handler chan []byte) {
				handler <- response
			}(handler)
		}
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
	options.SetConnectRetryInterval(time.Duration(c.broker.ConnectRetryDelay) * time.Second)

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
					c.logEntry.Info(message)
					c.emit(strconv.FormatInt(subscription.SubID, 10), message.Payload())
				}(message)
			}); token.Wait() && token.Error() != nil {
				c.logEntry.Error(token.Error())
			}
		default:
			continue
		}
	}
}
