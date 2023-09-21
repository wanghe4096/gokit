package gokit

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(data interface{}) {
	bytes, _ := json.Marshal(data)
	fmt.Println(string(bytes))
}
