module github.com/mysayasan/mqtt-client-go

go 1.16

require (
	github.com/eclipse/paho.mqtt.golang v1.3.4
	github.com/google/uuid v1.2.0
	github.com/mysayasan/helpers-go v0.0.0
	github.com/sirupsen/logrus v1.8.1
)

replace github.com/mysayasan/helpers-go v0.0.0 => ../helpers-go
