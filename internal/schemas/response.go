package schemas

import (
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Meta    interface{} `json:"meta"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(c *gin.Context, status int, message string, data interface{}, meta ...interface{}) {
	metaData := make(map[string]interface{})
	if len(meta) > 0 {
		metaData = meta[0].(map[string]interface{})
	}
	c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    metaData,
	})
}

func ErrorResponse(c *gin.Context, status int, message string) {
	pc, _, _, _ := runtime.Caller(1)
	funcName := runtime.FuncForPC(pc).Name()
	// Extract the function name from the full path
	parts := strings.Split(funcName, "/")
	shortFuncName := parts[len(parts)-1]
	fullMessage := shortFuncName + ": " + message
	c.JSON(status, Response{
		Status:  status,
		Message: fullMessage,
		Data:    nil,
	})
}

func ValidationErrorResponse(c *gin.Context, errors map[string]string) {
	c.JSON(http.StatusBadRequest, Response{
		Status:  http.StatusBadRequest,
		Message: "Validation error",
		Data:    errors,
	})
}
