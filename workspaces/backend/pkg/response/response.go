package response

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

type TResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type TDataResponse struct {
	TResponse
	Data interface{} `json:"data"`
}

type TErrResponse struct {
	TResponse
	Errors interface{} `json:"errors"`
}

func MessageResponse(c *gin.Context, code int) {
	c.JSON(getHttpCode(code), TResponse{
		Code:    code,
		Message: CodeMsg[code],
	})
}

func OkResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, TDataResponse{
		TResponse: TResponse{
			Code:    code,
			Message: CodeMsg[code],
		},
		Data: data,
	})
}

func CreatedResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusCreated, TDataResponse{
		TResponse: TResponse{
			Code:    code,
			Message: CodeMsg[code],
		},
		Data: data,
	})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(getHttpCode(code), TErrResponse{
		TResponse: TResponse{
			Code:    code,
			Message: CodeMsg[code],
		},
		Errors: err.Error(),
	})
}

func ValidateErrorResponse(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, TErrResponse{
		TResponse: TResponse{
			Code:    ErrCodeParamInvalid,
			Message: CodeMsg[ErrCodeParamInvalid],
		},
		Errors: validationErrorsToJSON(err),
	})
}

func getHttpCode(code int) int {
	switch true {
	case code >= 20000 && code < 20100:
		return http.StatusOK
	case code >= 20100 && code < 20200:
		return http.StatusOK
	case code >= 40000 && code < 40100:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func validationErrorsToJSON(err error) []string {
	var res []string

	// Check if it's a validation error
	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		for _, fieldErr := range validationErrs {
			// Collect error messages into a list
			if fieldErr.Tag() == "required" {
				res = append(res, fmt.Sprintf("Field %s is required", fieldErr.Field()))
			} else {
				res = append(res, fmt.Sprintf("Fiels %s is invalid due to %s", fieldErr.Field(), fieldErr.Tag()))
			}
		}
	}

	return res
}
