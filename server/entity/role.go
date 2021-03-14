package entity

import "time"

type Role struct {
	Key        string                            `json:"key" bson:"key"`
	Name       string                            `json:"Name" bson:"Name"`
	GroupPower map[string]map[string]interface{} `json:"group_power" bson:"group_power"` //config_create,config_list,config_update,config_delete,data_create,data_list,data_update,data_delete
	Creator    Creator                           `bson:"creator" json:"creator"`
	Updater    Updater                           `bson:"updater" json:"updater"`
}

// SetCreator .
func (m *Role) SetCreator(key string) {
	m.Creator.Key = key
	m.Updater.Key = key
	m.Creator.Timestamp = time.Now().UnixNano() / millisecond
	m.Updater.Timestamp = m.Creator.Timestamp
}

// SetUpdater .
func (m *Role) SetUpdater(key string) {
	m.Updater.Key = key
	m.Updater.Timestamp = time.Now().UnixNano() / millisecond
}
