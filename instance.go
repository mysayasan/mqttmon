package mqttmon

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	hosthelper "github.com/mysayasan/hosthelper"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
)

// instance struct
type instance struct {
	hub      *Hub
	brokers  []*Broker
	timeout  time.Duration
	logEntry *logrus.Entry
	// response chan interface{}
}

// NewInstance create new MQTT instancelication
// func NewSubscribe(hub *Hub, brokerUcase IBrokerUcase, subscriptionUcase ISubscribeUcase, timeout time.Duration, logger *logrus.Logger, response chan interface{}) ISubscribeInstance {
func NewInstance(hub *Hub, brokers []*Broker, timeout time.Duration, logger *logrus.Logger) IInstance {

	hostName, _ := os.Hostname()
	hostAddress, _ := hosthelper.ExternalIP()

	// Prepare logger
	logEntry := logger.WithFields(logrus.Fields{
		"eventid":     BrokerEvent,
		"hostname":    hostName,
		"hostaddress": hostAddress,
	})

	if len(brokers) < 1 {
		logEntry.Error(errors.New("no broker found"))
	}

	return &instance{
		hub:      hub,
		brokers:  brokers,
		timeout:  timeout,
		logEntry: logEntry,
		// response: make(chan interface{}),
	}
}

// Connect to MQTT Client
func (a *instance) Connect() error {
	if len(a.brokers) < 1 {
		return errors.New("no broker found")
	}
	// Set error log for mqtt
	mqtt.ERROR = log.New(a.logEntry.WriterLevel(2), "Error: ", log.LstdFlags)

	for _, broker := range a.brokers {
		// Create client
		// client := NewClient(broker, logEntry, onMessageReceived)
		client := NewClient(broker, a.logEntry)
		a.hub.Register <- client

		// Connect client to broker
		go client.Connect()
	}

	return nil
}

func (a *instance) MultiSubscribe(subscriptions []*Subscription) (chan []byte, error) {
	if len(a.brokers) < 1 {
		return nil, errors.New("no broker found")
	}

	if len(subscriptions) < 1 {
		return nil, errors.New("no subscriptions")
	}

	response := make(chan []byte)

	for _, subscription := range subscriptions {
		a.Subscribe(subscription, response)
	}

	return response, nil
}

func (a *instance) PublishChannel(publishChan chan []byte) {
	go a.publishChannel(publishChan)
}

func (a *instance) publishChannel(publishChan chan []byte) {
	for {
		select {
		case data := <-publishChan:
			var publication Publication
			err := json.Unmarshal(data, &publication)
			if err != nil {
				a.logEntry.Error(err)
				continue
			}
			// fmt.Printf("%v", publish)
			a.Publish(publication)
		default:
			continue
		}
	}
}

func (a *instance) Publish(publication Publication) {
	publishJSON, err := json.Marshal(publication)
	if err != nil {
		a.logEntry.Error(err)
		return
	}
	// Publish to mqtt client
	for client := range a.hub.Clients {
		// if client.broker.BrokerID == publication.BrokerID {
		// 	client.Publish <- publishJSON
		// }
		if client.broker.BrokerID == publication.BrokerID {
			client.Publish <- publishJSON
		}
	}
}

func (a *instance) Subscribe(subscription *Subscription, response chan []byte) {
	subscriptionJSON, err := json.Marshal(subscription)
	if err != nil {
		a.logEntry.Error(err)
		return
	}
	// Subscribe to mqtt broker
	for client := range a.hub.Clients {
		// if client.broker.BrokerID == subscription.BrokerID {
		// 	client.AddListener(strconv.FormatInt(subscription.SessionID, 10), response)
		// 	client.Subscribe <- subscriptionJSON
		// }
		if client.broker.BrokerID == subscription.BrokerID {
			// client.AddListener(strconv.FormatInt(subscription.SessionID, 10), response)
			client.AddListener(subscription.SessionID, response)
			client.Subscribe <- subscriptionJSON
		}
	}
}
