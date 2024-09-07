package utils

import (
	"encoding/json"
	"github.com/go-open-auth/global"
	"go.uber.org/zap"
	"reflect"
)

func ModelToDto[T any](model interface{}) interface{} {
	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	plain := make(map[string]interface{})
	var dto T
	dtoType := reflect.TypeOf(dto)

	for i := 0; i < modelType.NumField(); i++ {
		fieldType := modelType.Field(i)
		dtoFieldType, _ := dtoType.FieldByName(fieldType.Name)
		nestedKey := dtoFieldType.Tag.Get("nested")
		if nestedKey != "" {
			value := modelValue.Field(i)
			plain[fieldType.Name] = value.FieldByName(nestedKey).Interface()
		} else {
			plain[fieldType.Name] = modelValue.Field(i).Interface()
		}
	}

	bytes, _ := json.Marshal(plain)
	err := json.Unmarshal(bytes, &dto)
	if err != nil {
		global.Logger.Error("convert to dto failed", zap.Error(err))
		return nil
	}
	return dto
}
