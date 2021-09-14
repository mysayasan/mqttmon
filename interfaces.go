package mqttmon

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	RunSubscribe(subscriptions []*Subscription) (chan []byte, error)
	RunPublisher(publicationChan chan []byte)
	Publish(publication Publication)
	Subscribe(subscription *Subscription, response chan []byte)
}
