package entity

import "time"

// ConfigData .
type ConfigData map[string]interface{}

// SetCreator .
func (b ConfigData) SetCreator(id int64) {
	timestamp := time.Now().UnixNano() / millisecond

	creator := Creator{}
	creator.Set(id, timestamp)
	b["creator"] = creator

	updater := Updater{}
	updater.Set(id, timestamp)
	b["updater"] = updater
}

// SetUpdater .
func (b ConfigData) SetUpdater(id int64) {
	timestamp := time.Now().UnixNano() / millisecond

	updater := Updater{}
	updater.Set(id, timestamp)
	b["updater"] = updater

	//key和creator一旦定义里就禁止更新
	delete(b, "key")
	delete(b, "creator")
}
