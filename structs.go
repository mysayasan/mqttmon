package mqttmon

import (
	"time"

	"github.com/google/uuid"
)

// Broker Model
type Broker struct {
	BrokerID          uuid.UUID `json:"brokerid" form:"brokerid" query:"brokerid"`
	BrokerDescription string    `json:"brokerdescription" form:"brokerdescription" query:"brokerdescription"`
	BrokerAddress     string    `json:"brokeraddress" form:"brokeraddress" query:"brokeraddress" validate:"required"`
	ClientID          string    `json:"clientid" form:"clientid" query:"clientid"`
	Username          string    `json:"username" form:"username" query:"username"`
	Userpass          string    `json:"userpass" form:"userpass" query:"userpass"`
	PingTimeout       int       `json:"pingtimeout" form:"pingtimeout" query:"pingtimeout"`
	KeepAlive         int       `json:"keepalive" form:"keepalive" query:"keepalive"`
	AutoReconnect     int       `json:"autoreconnect" form:"autoreconnect" query:"autoreconnect"`
	ConnectRetryDelay int       `json:"connectretrydelay" form:"connectretrydelay" query:"connectretrydelay"`
	IsActive          int       `json:"isactive" form:"isactive" query:"isactive"`
	CreatedBy         string    `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn         int64     `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated       time.Time `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy         string    `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn         int64     `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated       time.Time `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
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
	CreatedBy          string      `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn          int64       `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated        time.Time   `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy          string      `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn          int64       `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated        time.Time   `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
}

// Subscription Model
type Subscription struct {
	SubscriptionID          uuid.UUID `json:"subscriptionid" form:"subscriptionid" query:"subscriptionid"`
	SubscriptionDescription string    `json:"subscriptiondescription" form:"subscriptiondescription" query:"subscriptiondescription"`
	BrokerID                uuid.UUID `json:"brokerid" form:"brokerid" query:"brokerid" validate:"required"`
	Topic                   string    `json:"topic" form:"topic" query:"topic"`
	QOS                     int       `json:"qos" form:"qos" query:"qos"`
	IsActive                int       `json:"isactive" form:"isactive" query:"isactive"`
	CreatedBy               string    `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn               int64     `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated             time.Time `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy               string    `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn               int64     `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated             time.Time `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
}
