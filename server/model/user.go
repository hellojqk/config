package model

type UserLoginParam struct {
	Key      string `json:"key" form:"key" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}
