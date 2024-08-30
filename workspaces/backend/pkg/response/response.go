package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  interface{} `json:"errors"`
}

func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: ErrCodeMsg[code],
		Data:    data,
	})
}

func ErrorResponse(c *gin.Context, code int, err error) {
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: ErrCodeMsg[code],
		Data:    nil,
		Errors:  err.Error(),
	})
}
