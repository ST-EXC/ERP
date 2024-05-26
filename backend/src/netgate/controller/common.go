package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type JsonStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}
type JsonErrorStruct struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func Success(c *gin.Context, code int, message interface{}, data interface{}) {
	json := &JsonStruct{Code: code, Message: message, Data: data}
	c.JSON(http.StatusAccepted, json)
}

func Error(c *gin.Context, code int, message interface{}) {
	json := &JsonErrorStruct{Code: code, Message: message}
	c.JSON(http.StatusBadRequest, json)
}
