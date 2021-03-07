package entity

import "go.mongodb.org/mongo-driver/bson"

// ConfigData .
type ConfigData struct {
	Key  string `json:"key"`
	Data bson.M `json:"data"`
	Base
}
