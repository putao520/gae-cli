package gsc

import (
	"encoding/json"
)

func MapToStruct[T any](body map[string]interface{}) *T {
	b, err := json.Marshal(body)
	if err != nil {
		return nil
	}
	var t T
	err = json.Unmarshal(b, &t)
	if err != nil {
		return nil
	}
	return &t
}
