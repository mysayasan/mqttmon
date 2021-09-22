package mqttmon

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	MultiSubscribe(subscriptions []*Subscription) (chan []byte, error)
	StreamPublish(publicationChan chan []byte)
	Publish(publication Publication)
	Subscribe(subscription *Subscription, response chan []byte)
}
