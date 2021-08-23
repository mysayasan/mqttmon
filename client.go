package mqttclient

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

// // responseMessage standard modbus response format
// type responseMessage struct {
// 	TimeStamp int64  `json:"timestamp" form:"timestamp" query:"timestamp"`
// 	Topic     string `json:"topic" form:"topic" query:"topic"`
// 	Payload   []byte `json:"payload" form:"payload" query:"payload"`
// }

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

	// listeners
	// listeners map[string]chan []byte
	listeners map[string][]chan []byte

	// // Event listeners
	// Event event.Event

	// Message handlers
	// messageHandler mqtt.MessageHandler
	// messageHandler MessageHandler
}

// NewClient create new client
func NewClient(
	broker *Broker,
	logEntry *logrus.Entry,
	// messageHandler mqtt.MessageHandler,
	// messageHandler MessageHandler,
) *Client {
	return &Client{
		broker:    broker,
		logEntry:  logEntry,
		Publish:   make(chan []byte),
		Subscribe: make(chan []byte),
		Close:     make(chan bool),
		// messageHandler: messageHandler,
	}
}

// // AddListener add event listener
// func (c *Client) AddListener(e string, ch chan []byte) {
// 	if c.listeners == nil {
// 		c.listeners = make(map[string]chan []byte)
// 	}
// 	if _, ok := c.listeners[e]; ok {
// 		close(c.listeners[e])
// 		c.listeners[e] = ch
// 	} else {
// 		c.listeners[e] = ch
// 	}
// }

// // RemoveListener removes an event listener
// func (c *Client) RemoveListener(e string, ch chan []byte) {
// 	if _, ok := c.listeners[e]; ok {
// 		delete(c.listeners, e)
// 		close(c.listeners[e])
// 	}
// }

// // Emit emits an event on the Dog struct instance
// func (c *Client) emit(e string, response []byte) {
// 	if _, ok := c.listeners[e]; ok {
// 		c.listeners[e] <- response
// 		// for _, handler := range c.listeners {
// 		// 	go func(handler chan []byte) {
// 		// 		handler <- response
// 		// 	}(handler)
// 		// }
// 	}
// }

// AddListener add event listener
func (c *Client) AddListener(e string, ch chan []byte) {
	if c.listeners == nil {
		c.listeners = make(map[string][]chan []byte)
	}

	c.listeners[e] = append(c.listeners[e], ch)

	// if _, ok := c.listeners[e]; ok {
	// 	c.listeners[e] = append(c.listeners[e], ch)
	// } else {
	// 	c.listeners[e] = []chan []byte{ch}
	// }
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

// // onMessageReceived Print out message
// func (c *Client) onMessageReceived(client mqtt.Client, message mqtt.Message) {
// 	// messageJSON, err := json.Marshal(message.Payload())
// 	// if err != nil {
// 	// 	c.logEntry.Error(err)
// 	// }
// 	c.emit("test", message.Payload())
// 	// fmt.Println(message.Payload())
// }

// Connect to  Broker
// func (c *Client) Connect(wg *sync.WaitGroup) {
func (c *Client) Connect() {
	hostName, _ := os.Hostname()

	var clientid = ""
	if c.broker.ClientID == "nil" || c.broker.ClientID == "" {
		clientid = hostName + strconv.Itoa(time.Now().Second())
	}

	options := mqtt.NewClientOptions().AddBroker(c.broker.BrokerAddress).SetClientID(clientid).SetCleanSession(true)
	if c.broker.ClientUsername != "" && c.broker.ClientID != "nil" {
		options.SetUsername(c.broker.ClientUsername)
		if c.broker.ClientPassword != "" && c.broker.ClientPassword != "nil" {
			options.SetPassword(c.broker.ClientPassword)
		}
	}

	tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	options.SetTLSConfig(tlsConfig)
	options.SetAutoReconnect(c.broker.AutoReconnect == 1)
	options.SetKeepAlive(time.Duration(c.broker.KeepAlive) * time.Second)
	options.SetPingTimeout(time.Duration(c.broker.PingTimeout) * time.Second)
	options.SetConnectRetryInterval(time.Duration(c.broker.ConnectRetryInterval) * time.Second)

	// // Test Subscribe
	// options.OnConnect = func(mc mqtt.Client) {
	// 	if token := mc.Subscribe("serverroom/#", byte(0), c.onMessageReceived); token.Wait() && token.Error() != nil {
	// 		c.logEntry.Error(token.Error())
	// 	}
	// }

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

			var publish Publish

			err := json.Unmarshal(data, &publish)
			if err != nil {
				c.logEntry.Error(err)
				continue
			}

			messageJSON, err := json.Marshal(publish.Payload)
			if err != nil {
				c.logEntry.Error(err)
			}

			// fmt.Println(publish)

			if token := client.Publish(publish.Topic, byte(publish.QOS), publish.IsRetain == 1, messageJSON); token.Wait() && token.Error() != nil {
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
			// if token := client.Subscribe(subscription.Topic, byte(subscription.QOS), c.onMessageReceived); token.Wait() && token.Error() != nil {
			// 	c.logEntry.Error(token.Error())
			// }

			if token := client.Subscribe(subscription.Topic, byte(subscription.QOS), func(client mqtt.Client, message mqtt.Message) {
				go func(message mqtt.Message) {
					// res := responseMessage{
					// 	TimeStamp: time.Now().Local().Unix(),
					// 	Topic:     message.Topic(),
					// 	Payload:   message.Payload(),
					// }

					// resJSON, err := json.Marshal(res.Payload)
					// payload := string(message.Payload())
					// fmt.Println(payload)
					// if err != nil {
					// 	c.logEntry.Error(err)
					// }
					// fmt.Printf("Timestamp: %d, %s\n", time.Now().Unix(), resJSON)

					c.logEntry.Info(message)
					c.emit(subscription.SubscriptionID.String(), message.Payload())
				}(message)

				// res := responseMessage{
				// 	TimeStamp: time.Now().Local().Unix(),
				// 	Topic:     message.Topic(),
				// 	Payload:   string(message.Payload()),
				// }
				// resJSON, err := json.Marshal(res)
				// if err != nil {
				// 	c.logEntry.Error(err)
				// }
				// c.emit(subscription.SubscriptionID.String(), message)
			}); token.Wait() && token.Error() != nil {
				c.logEntry.Error(token.Error())
			}
		default:
			continue
		}
	}
}
