package entity

import "time"

// ConfigStruct .
type ConfigStruct struct {
	Key         string  `bson:"key" json:"key"`
	Title       string  `bson:"title" json:"title"`
	Description string  `bson:"description" json:"description"`
	Secret      bool    `bson:"secret" json:"secret"`
	Array       bool    `bson:"array" json:"array"`
	Schema      string  `bson:"schema" json:"schema"`
	Creator     Creator `bson:"creator" json:"creator"`
	Updater     Updater `bson:"updater" json:"updater"`
}

// SetCreator .
func (m *ConfigStruct) SetCreator(id int64) {
	m.Creator.ID = id
	m.Updater.ID = id
	m.Creator.Timestamp = time.Now().UnixNano() / millisecond
	m.Updater.Timestamp = m.Creator.Timestamp
}

// SetUpdater .
func (m *ConfigStruct) SetUpdater(id int64) {
	m.Updater.ID = id
	m.Updater.Timestamp = time.Now().UnixNano() / millisecond
}
