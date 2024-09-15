package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/open-auth/global"
	"github.com/open-auth/internal/dto"
	"github.com/open-auth/pkg/response"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func ModelToDto[T any, MT any](model MT) *T {
	modelType := reflect.TypeOf(model)
	modelValue := reflect.ValueOf(model)
	plain := make(map[string]interface{})
	var payload T
	dtoType := reflect.TypeOf(payload)

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
	err := json.Unmarshal(bytes, &payload)
	if err != nil {
		global.Logger.Error("convert to dto failed", zap.Error(err))
		return nil
	}
	return &payload
}

func ModelToDtos[T any, MT any](models []MT) []T {
	list := make([]T, len(models))
	for i := range list {
		list[i] = *ModelToDto[T, MT](models[i])
	}
	return list
}

func ModelToPaginationDto[T any, MT any](models []MT, metaData dto.PaginationMetaDataDto) dto.PaginationDto[T] {
	list := ModelToDtos[T, MT](models)
	return dto.PaginationDto[T]{
		Data:     list,
		MetaData: metaData,
	}
}

func BodyToDto[T any](c *gin.Context) *T {
	var payload *T
	if err := c.ShouldBindBodyWithJSON(&payload); err != nil {
		response.ValidateErrorResponse(c, err)
		c.Abort()
		return nil
	}
	return payload
}

func QueryToDto[T any](c *gin.Context) T {
	var payload T
	if err := c.ShouldBindQuery(&payload); err != nil {
		global.Logger.Error("convert to dto failed", zap.Error(err))
		response.ValidateErrorResponse(c, err)
		c.Abort()
	}
	return payload
}

func DtoToModel[MT any, T any](dto T) (*MT, *response.ServerCode) {
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
		return nil, response.ReturnCode(response.ErrCodeParamInvalid)
	}

	return &model, nil
}

func PartialUpdate[T any](payload T) string {
	payloadValue := reflect.ValueOf(payload)
	payloadType := reflect.TypeOf(payload)
	var arr []string
	for i := 0; i < payloadType.NumField(); i++ {
		field := payloadType.Field(i)
		fieldName := field.Name
		payloadFieldValue := payloadValue.FieldByName(fieldName)
		dbName := field.Tag.Get("db")
		attrName := field.Tag.Get("attr")
		if payloadFieldValue.IsNil() != true && attrName != "" {
			arr = append(arr, fmt.Sprintf("%s = :%s", attrName, dbName))
		}
	}
	return strings.Join(arr, ", ")
}
