package utils

import "encoding/json"

func IsEmpty(s string) bool {
	return s == ""
}

func SafeJson(a interface{}) string {
	v, err := json.Marshal(a)
	if err != nil {
		return "{}"
	}
	return string(v)
}
