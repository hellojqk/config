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
	ID        int64 `bson:"id,omitempty" json:"id"`
	Timestamp int64 `bson:"timestamp,omitempty" json:"timestamp"`
}

// Set .
func (m *Creator) Set(id int64, timestamp int64) {
	m.ID = id
	m.Timestamp = timestamp
}

// Updater .
type Updater struct {
	ID        int64 `bson:"id,omitempty" json:"id"`
	Timestamp int64 `bson:"timestamp,omitempty" json:"timestamp"`
}

var millisecond = int64(time.Millisecond)

// Set .
func (m *Updater) Set(id int64, timestamp int64) {
	m.ID = id
	m.Timestamp = timestamp
}
