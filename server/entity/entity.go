package entity

import "time"

// ListPagingParam .
type ListPagingParam struct {
	PageNum  int64       `json:"page_num" form:"page_num" query:"page_num"`
	PageSize int64       `json:"page_size" form:"page_size" query:"page_size"`
	Filter   interface{} `json:"filter" form:"filter" query:"filter"`
	Sort     interface{} `json:"sort" form:"sort" query:"sort"`
}

// Creator .
type Creator struct {
	Key       string `bson:"key,omitempty" json:"key"`
	Timestamp int64  `bson:"timestamp,omitempty" json:"timestamp"`
}

// Set .
func (m *Creator) Set(key string, timestamp int64) {
	m.Key = key
	m.Timestamp = timestamp
}

// Updater .
type Updater struct {
	Key       string `bson:"key,omitempty" json:"key"`
	Timestamp int64  `bson:"timestamp,omitempty" json:"timestamp"`
}

var millisecond = int64(time.Millisecond)

// Set .
func (m *Updater) Set(key string, timestamp int64) {
	m.Key = key
	m.Timestamp = timestamp
}
