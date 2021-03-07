package util

import (
	"encoding/json"
	"fmt"
)

// PrintJSON .
func PrintJSON(str string, i interface{}) {
	bts, _ := json.Marshal(i)
	fmt.Printf("%s：%s\n", str, bts)
}
