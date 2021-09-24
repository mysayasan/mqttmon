package mqttmon

// IInstance represent MQTT application
type IInstance interface {
	Connect() error
	SubscribeAll(subscriptions []*Subscription) (chan []byte, error)
	PublishChannel(publicationChan chan []byte)
	Publish(publication Publication) error
	Subscribe(subscription *Subscription, response chan []byte) error
	Unsubscribe(subscription *Subscription) error
	UnsubscribeAll(subscription []*Subscription) error
}
