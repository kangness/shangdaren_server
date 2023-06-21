package model

// SDRRequest ...
type SDRRequest struct {
	RequestId string      `json:"request_id,omitempty"`
	ActionCmd string      `json:"action_cmd,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// SDRGetCounterRequest ..
type SDRGetCounterRequest struct {
	Id  int64  `json:"id,omitempty"`
	Ext string `json:"ext,omitempty"`
}

// SDRSetCounterRequest ...
type SDRSetCounterRequest struct {
	Id     int64  `json:"id,omitempty"`
	Action string `json:"action,omitempty"`
	Value  int64  `json:"value,omitempty"`
}

// SDRSetCounterResponse ...
type SDRSetCounterResponse struct {
	Id         int64  `json:"id,omitempty"`
	Count      int64  `json:"count,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Ext        string `json:"ext,omitempty"`
}

// SDRGetCounterResponse ...
type SDRGetCounterResponse struct {
	Id         int64  `json:"id,omitempty"`
	Count      int64  `json:"count,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	UpdateTime string `json:"update_time,omitempty"`
	Ext        string `json:"ext,omitempty"`
}

// SDRResponse ...
type SDRResponse struct {
	RequestId string      `json:"request_id,omitempty"`
	Code      int64       `json:"code,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}

// TCounters ...
type TCounters struct {
	Id        int64  `gorm:"column:id" db:"id" json:"id" form:"id"`
	Count     int64  `gorm:"column:count" db:"count" json:"count" form:"count"`
	Createdat string `gorm:"column:createdAt" db:"createdAt" json:"createdAt" form:"createdAt"`
	Updatedat string `gorm:"column:updatedAt" db:"updatedAt" json:"updatedAt" form:"updatedAt"`
}

// TableName ...
func (*TCounters) TableName() string {
	return "TCounters"
}
