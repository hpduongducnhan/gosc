package utils

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
)

type StructContraint interface {
}

func JsonString2Struct[T StructContraint](jsonString string) (result T, err error) {
	var rawMap map[string]interface{}
	err = json.Unmarshal([]byte(jsonString), &rawMap)
	if err != nil {
		return
	}

	err = mapstructure.Decode(rawMap, &result)
	if err != nil {
		return
	}
	return
}
