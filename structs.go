package mqttmon

import (
	"time"

	"github.com/google/uuid"
)

// Broker Model
type Broker struct {
	BrokerID             uuid.UUID `json:"brokerid" form:"brokerid" query:"brokerid"`
	BrokerDescription    string    `json:"brokerdescription" form:"brokerdescription" query:"brokerdescription"`
	BrokerAddress        string    `json:"brokeraddress" form:"brokeraddress" query:"brokeraddress" validate:"required"`
	ClientID             string    `json:"clientid" form:"clientid" query:"clientid"`
	Username             string    `json:"username" form:"username" query:"username"`
	Userpass             string    `json:"userpass" form:"userpass" query:"userpass"`
	PingTimeout          int       `json:"pingtimeout" form:"pingtimeout" query:"pingtimeout"`
	KeepAlive            int       `json:"keepalive" form:"keepalive" query:"keepalive"`
	AutoReconnect        int       `json:"autoreconnect" form:"autoreconnect" query:"autoreconnect"`
	ConnectRetryInterval int       `json:"connectretryinterval" form:"connectretryinterval" query:"connectretryinterval"`
	IsActive             int       `json:"isactive" form:"isactive" query:"isactive"`
	CreatedUserName      string    `json:"createdusername" form:"createdusername" query:"createdusername"`
	CreatedDateUnix      int64     `json:"createddateunix" form:"createddateunix" query:"createddateunix"`
	CreatedDate          time.Time `json:"createddate" form:"createddate" query:"createddate"`
	ModifiedUserName     string    `json:"modifiedusername" form:"modifiedusername" query:"modifiedusername"`
	ModifiedDateUnix     int64     `json:"modifieddateunix" form:"modifieddateunix" query:"modifieddateunix"`
	ModifiedDate         time.Time `json:"modifydate" form:"modifydate" query:"modifydate"`
}

// Publish Model
type Publish struct {
	PublishID          uuid.UUID   `json:"publishid" form:"publishid" query:"publishid"`
	PublishDescription string      `json:"publishdescription" form:"publishdescription" query:"publishdescription"`
	BrokerID           uuid.UUID   `json:"brokerid" form:"brokerid" query:"brokerid" validate:"required"`
	Topic              string      `json:"topic" form:"topic" query:"topic"`
	QOS                int         `json:"qos" form:"qos" query:"qos"`
	IsRetain           int         `json:"isretain" form:"isretain" query:"isretain"`
	Payload            interface{} `json:"payload" form:"payload" query:"payload"`
	IsActive           int         `json:"isactive" form:"isactive" query:"isactive"`
	CreatedUserName    string      `json:"createdusername" form:"createdusername" query:"createdusername"`
	CreatedDateUnix    int64       `json:"createddateunix" form:"createddateunix" query:"createddateunix"`
	CreatedDate        time.Time   `json:"createddate" form:"createddate" query:"createddate"`
	ModifiedUserName   string      `json:"modifiedusername" form:"modifiedusername" query:"modifiedusername"`
	ModifiedDateUnix   int64       `json:"modifieddateunix" form:"modifieddateunix" query:"modifieddateunix"`
	ModifiedDate       time.Time   `json:"modifydate" form:"modifydate" query:"modifydate"`
}

// Subscription Model
type Subscription struct {
	SubscriptionID          uuid.UUID `json:"subscriptionid" form:"subscriptionid" query:"subscriptionid"`
	SubscriptionDescription string    `json:"subscriptiondescription" form:"subscriptiondescription" query:"subscriptiondescription"`
	BrokerID                uuid.UUID `json:"brokerid" form:"brokerid" query:"brokerid" validate:"required"`
	Topic                   string    `json:"topic" form:"topic" query:"topic"`
	QOS                     int       `json:"qos" form:"qos" query:"qos"`
	IsActive                int       `json:"isactive" form:"isactive" query:"isactive"`
	CreatedUserName         string    `json:"createdusername" form:"createdusername" query:"createdusername"`
	CreatedDateUnix         int64     `json:"createddateunix" form:"createddateunix" query:"createddateunix"`
	CreatedDate             time.Time `json:"createddate" form:"createddate" query:"createddate"`
	ModifiedUserName        string    `json:"modifiedusername" form:"modifiedusername" query:"modifiedusername"`
	ModifiedDateUnix        int64     `json:"modifieddateunix" form:"modifieddateunix" query:"modifieddateunix"`
	ModifiedDate            time.Time `json:"modifydate" form:"modifydate" query:"modifydate"`
}
