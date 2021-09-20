package mqttmon

import "time"

// Broker Model
type Broker struct {
	BrokerID       int       `json:"brokerid" form:"brokerid" query:"brokerid"`
	BrokerName     string    `json:"brokername" form:"brokername" query:"brokername"`
	BrokerAddress  string    `json:"brokeraddress" form:"brokeraddress" query:"brokeraddress" validate:"required"`
	ClientID       string    `json:"clientid" form:"clientid" query:"clientid"`
	Username       string    `json:"username" form:"username" query:"username"`
	Userpass       string    `json:"userpass" form:"userpass" query:"userpass"`
	PingTimeout    int       `json:"pingtimeout" form:"pingtimeout" query:"pingtimeout"`
	KeepAlive      int16     `json:"keepalive" form:"keepalive" query:"keepalive"`
	AutoReconnect  int16     `json:"autoreconnect" form:"autoreconnect" query:"autoreconnect"`
	ConnRetryDelay int       `json:"connretrydelay" form:"connretrydelay" query:"connretrydelay"`
	IsActive       int16     `json:"isactive" form:"isactive" query:"isactive"`
	CreatedBy      string    `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn      int64     `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated    time.Time `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy      string    `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn      int64     `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated    time.Time `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
}

// Publication Model
type Publication struct {
	PubID       int64       `json:"pubid" form:"pubid" query:"pubid"`
	PubDesc     string      `json:"pubdesc" form:"pubdesc" query:"pubdesc" validate:"required"`
	Topic       string      `json:"topic" form:"topic" query:"topic"`
	QOS         int16       `json:"qos" form:"qos" query:"qos"`
	IsRetain    int16       `json:"isretain" form:"isretain" query:"isretain"`
	Payload     interface{} `json:"payload" form:"payload" query:"payload"`
	IsActive    int16       `json:"isactive" form:"isactive" query:"isactive"`
	CreatedBy   string      `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn   int64       `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated time.Time   `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy   string      `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn   int64       `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated time.Time   `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
}

// Subscription Model
type Subscription struct {
	SubID       int64     `json:"subid" form:"subid" query:"subid"`
	SubDesc     string    `json:"subdesc" form:"subdesc" query:"subdesc" validate:"required"`
	Topic       string    `json:"topic" form:"topic" query:"topic"`
	QOS         int16     `json:"qos" form:"qos" query:"qos"`
	IsActive    int16     `json:"isactive" form:"isactive" query:"isactive"`
	CreatedBy   string    `json:"createdby" form:"createdby" query:"createdby"`
	CreatedOn   int64     `json:"createdon" form:"createdon" query:"createdon"`
	DateCreated time.Time `json:"datecreated" form:"datecreated" query:"datecreated"`
	UpdatedBy   string    `json:"updatedby" form:"updatedby" query:"updatedby"`
	UpdatedOn   int64     `json:"updatedon" form:"updatedon" query:"updatedon"`
	DateUpdated time.Time `json:"dateupdated" form:"dateupdated" query:"dateupdated"`
}
