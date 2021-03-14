package entity

import "time"

type User struct {
	Key        string                            `json:"key" bson:"key"`
	Name       string                            `json:"Name" bson:"Name"`
	RoleKeys   []string                          `json:"role_keys" bson:"role_keys"`
	Password   string                            `json:"-" bson:"password"`
	GroupPower map[string]map[string]interface{} `json:"group_power" bson:"group_power"`
	Creator    Creator                           `bson:"creator" json:"creator"`
	Updater    Updater                           `bson:"updater" json:"updater"`
}

// SetCreator .
func (m *User) SetCreator(key string) {
	m.Creator.Key = key
	m.Updater.Key = key
	m.Creator.Timestamp = time.Now().UnixNano() / millisecond
	m.Updater.Timestamp = m.Creator.Timestamp
}

// SetUpdater .
func (m *User) SetUpdater(key string) {
	m.Updater.Key = key
	m.Updater.Timestamp = time.Now().UnixNano() / millisecond
}
