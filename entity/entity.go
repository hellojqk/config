package entity

import "time"

// ListPagingParam .
type ListPagingParam struct {
	PageNum  int64       `json:"page_num" form:"page_num" query:"page_num"`
	PageSize int64       `json:"page_size" form:"page_size" query:"page_size"`
	Filter   interface{} `json:"filter" form:"filter" query:"filter"`
	Sort     interface{} `json:"sort" form:"sort" query:"sort"`
}

// Base .
type Base struct {
	CreateUserID int64 `json:"create_user_id"`
	CreateTime   int64 `json:"create_time"`
	UpdateUserID int64 `json:"update_user_id"`
	UpdateTime   int64 `json:"update_time"`
}

var millisecond = int64(time.Millisecond)

// Create .
func (b *Base) Create(userID int64) {
	b.CreateUserID = userID
	b.CreateTime = time.Now().UnixNano() / millisecond
}

// Update .
func (b *Base) Update(userID int64) {
	b.UpdateTime = userID
	b.UpdateTime = time.Now().UnixNano() / millisecond
}
