package util

import (
	"encoding/json"
	"io/ioutil"
)

func LoadJsonToMap(filename string) map[string]interface{} {
	data, _ := ioutil.ReadFile(filename)
	dataJson := make(map[string]interface{})
	json.Unmarshal(data, &dataJson)
	return dataJson
}

func LoadJsonToStruct(filename string, v interface{}) error {
	data, _ := ioutil.ReadFile(filename)
	return json.Unmarshal(data, &v)
}
