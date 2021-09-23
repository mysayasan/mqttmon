package mqttmon

// Broker Model
type Broker struct {
	BrokerID       int    `json:"brokerid" form:"brokerid" query:"brokerid"`
	BrokerName     string `json:"brokername" form:"brokername" query:"brokername"`
	BrokerAddress  string `json:"brokeraddress" form:"brokeraddress" query:"brokeraddress" validate:"required"`
	ClientID       string `json:"clientid" form:"clientid" query:"clientid"`
	Username       string `json:"username" form:"username" query:"username"`
	Userpass       string `json:"userpass" form:"userpass" query:"userpass"`
	PingTimeout    int    `json:"pingtimeout" form:"pingtimeout" query:"pingtimeout"`
	KeepAlive      int16  `json:"keepalive" form:"keepalive" query:"keepalive"`
	AutoReconnect  int16  `json:"autoreconnect" form:"autoreconnect" query:"autoreconnect"`
	ConnRetryDelay int    `json:"connretrydelay" form:"connretrydelay" query:"connretrydelay"`
}

// Publication Model
type Publication struct {
	BrokerID int         `json:"brokerid" form:"brokerid" query:"brokerid" validate:"required"`
	PubDesc  string      `json:"pubdesc" form:"pubdesc" query:"pubdesc" validate:"required"`
	Topic    string      `json:"topic" form:"topic" query:"topic"`
	QOS      int16       `json:"qos" form:"qos" query:"qos"`
	IsRetain int16       `json:"isretain" form:"isretain" query:"isretain"`
	Payload  interface{} `json:"payload" form:"payload" query:"payload"`
}

// Subscription Model
type Subscription struct {
	SessionID string `json:"sessionid" form:"sessionid" query:"sessionid" validate:"required"`
	BrokerID  int    `json:"brokerid" form:"brokerid" query:"brokerid" validate:"required"`
	SubDesc   string `json:"subdesc" form:"subdesc" query:"subdesc" validate:"required"`
	Topic     string `json:"topic" form:"topic" query:"topic"`
	QOS       int16  `json:"qos" form:"qos" query:"qos"`
}
