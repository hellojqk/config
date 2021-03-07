package entity

// ConfigStruct .
type ConfigStruct struct {
	Key         string `json:"key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Secret      bool   `json:"secret"`
	Array       bool   `json:"array"`
	Schema      string `json:"schema"`
	Base
}
