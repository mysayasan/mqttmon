package mqttmon

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	MultiSubscribe(brokerid int, subscriptions []*Subscription) (chan []byte, error)
	PublishChannel(brokerid int, publicationChan chan []byte)
	Publish(brokerid int, publication Publication)
	Subscribe(brokerid int, subscription *Subscription, response chan []byte)
}
