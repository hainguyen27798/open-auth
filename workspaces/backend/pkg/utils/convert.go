package utils

import (
	"database/sql"
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

func DtoToModel[MT any, T any](dto T) (*MT, *int) {
	dtoType := reflect.TypeOf(dto)
	dtoValue := reflect.ValueOf(dto)
	plain := make(map[string]interface{})
	var model MT

	for i := 0; i < dtoType.NumField(); i++ {
		fieldType := dtoType.Field(i)
		mappingType := fieldType.Tag.Get("mappingType")
		if mappingType == "NullString" {
			fieldValue := dtoValue.Field(i)
			var value string
			if fieldValue.IsNil() == false {
				value = *fieldValue.Interface().(*string)
			}
			plain[fieldType.Name] = sql.NullString{
				String: value,
				Valid:  fieldValue.IsNil() == false,
			}
		} else {
			plain[fieldType.Name] = dtoValue.Field(i).Interface()
		}
	}

	bytes, _ := json.Marshal(plain)
	err := json.Unmarshal(bytes, &model)
	if err != nil {
		global.Logger.Error("convert to dto failed", zap.Error(err))
		return nil, &[]int{response.ErrCodeParamInvalid}[0]
	}

	return &model, nil
}
