package entity

import "time"

// ConfigStruct .
type ConfigStruct struct {
	Key         string  `bson:"key" json:"key" binding:"required"`
	GroupKey    string  `json:"group_key" bson:"group_key" binding:"required"`
	Title       string  `bson:"title" json:"title" binding:"required"`
	Description string  `bson:"description" json:"description"`
	Secret      bool    `bson:"secret" json:"secret"`
	Array       bool    `bson:"array" json:"array"`
	Schema      string  `bson:"schema" json:"schema" binding:"required"`
	Creator     Creator `bson:"creator" json:"creator"`
	Updater     Updater `bson:"updater" json:"updater"`
}

// SetCreator .
func (m *ConfigStruct) SetCreator(key string) {
	m.Creator.Key = key
	m.Updater.Key = key
	m.Creator.Timestamp = time.Now().UnixNano() / millisecond
	m.Updater.Timestamp = m.Creator.Timestamp
}

// SetUpdater .
func (m *ConfigStruct) SetUpdater(key string) {
	m.Updater.Key = key
	m.Updater.Timestamp = time.Now().UnixNano() / millisecond
}
