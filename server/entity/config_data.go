package entity

import "time"

// ConfigData .
type ConfigData map[string]interface{}

// SetCreator .
func (b ConfigData) SetCreator(key string) {
	timestamp := time.Now().UnixNano() / millisecond

	creator := Creator{}
	creator.Set(key, timestamp)
	b["creator"] = creator

	updater := Updater{}
	updater.Set(key, timestamp)
	b["updater"] = updater
}

// SetUpdater .
func (b ConfigData) SetUpdater(key string) {
	timestamp := time.Now().UnixNano() / millisecond

	updater := Updater{}
	updater.Set(key, timestamp)
	b["updater"] = updater

	//key和creator一旦定义里就禁止更新
	delete(b, "key")
	delete(b, "creator")
}
