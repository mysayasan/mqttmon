package mqttclient

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	RunSubscribe(subscriptions []*Subscription) (chan []byte, error)
	RunPublisher(publishChan chan []byte)
	Publish(publish Publish)
	Subscribe(subscription *Subscription, response chan []byte)
}
