package entity

import "time"

type Group struct {
	Key     string  `json:"key" bson:"key"`
	Name    string  `json:"Name" bson:"Name"`
	Creator Creator `bson:"creator" json:"creator"`
	Updater Updater `bson:"updater" json:"updater"`
}

// SetCreator .
func (m *Group) SetCreator(key string) {
	m.Creator.Key = key
	m.Updater.Key = key
	m.Creator.Timestamp = time.Now().UnixNano() / millisecond
	m.Updater.Timestamp = m.Creator.Timestamp
}

// SetUpdater .
func (m *Group) SetUpdater(key string) {
	m.Updater.Key = key
	m.Updater.Timestamp = time.Now().UnixNano() / millisecond
}
