package mqttmon

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	RunSubscription(brokerid int, subscriptions []*Subscription) (chan []byte, error)
	RunPublication(brokerid int, publicationChan chan []byte)
	Publish(brokerid int, publication Publication)
	Subscribe(brokerid int, subscription *Subscription, response chan []byte)
}
