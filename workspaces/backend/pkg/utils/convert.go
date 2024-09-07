package utils

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-open-auth/global"
	"github.com/go-open-auth/pkg/response"
	"go.uber.org/zap"
	"reflect"
)

func ModelToDto[T any, MT any](model MT) *T {
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
	return &dto
}

func ModelToDtos[T any, MT any](models []MT) []T {
	list := make([]T, len(models))
	for i := range list {
		list[i] = *ModelToDto[T, MT](models[i])
	}
	return list
}

func BodyToDto[T any](c *gin.Context) *T {
	var dto *T
	if err := c.ShouldBindBodyWithJSON(&dto); err != nil {
		response.ValidateErrorResponse(c, err)
		c.Abort()
		return nil
	}
	return dto
}
