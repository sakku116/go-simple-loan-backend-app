package helper

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func ArrayContains(arr interface{}, item interface{}) bool {
	newArr, ok := arr.([]interface{})
	if !ok {
		return false
	}

	for _, v := range newArr {
		if v == item {
			return true
		}
	}
	return false
}

func PrettyJson(data interface{}) string {
	res, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Sprintf("<failed to parse json: %v>", err.Error())
	}
	return string(res)
}

func GetStructAttributesJson(s interface{}, exclude []string, excludeJsonValue []string) []string {
	var attributes []string
	t := reflect.TypeOf(s)

	// Loop through the fields of the struct
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if exclude != nil {
			if ArrayContains(exclude, field.Name) {
				continue
			}
		}

		jsonValue := field.Tag.Get("json")
		if jsonValue == "" {
			continue
		}

		if excludeJsonValue != nil {
			if ArrayContains(excludeJsonValue, jsonValue) {
				continue
			}
		}
		attributes = append(attributes, jsonValue)
	}

	return attributes
}
