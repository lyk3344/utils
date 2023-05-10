package utils

import "encoding/json"

// json转Map ()
func JSONToMap(str string) map[string]interface{} {
	var tempMap = make(map[string]interface{})
	err := json.Unmarshal([]byte(str), &tempMap)
	if err != nil {
		panic(err)
	}
	return tempMap
}
