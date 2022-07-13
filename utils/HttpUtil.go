package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

var successMsg = "success" //成功默认值

func Success(context *gin.Context, data interface{}, msg ...string) {
	if len(msg) > 0 {
		successMsg = msg[0]
	}
	context.JSON(200, gin.H{
		"error":     0,
		"message":   successMsg,
		"data":      data,
		"timestamp": time.Now().Unix(),
	})
}

func Error(context *gin.Context, msg interface{}) {
	context.JSON(400, gin.H{
		"error":     1,
		"message":   msg,
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
}

//表单验证错误返回
func ValidateError(context *gin.Context, msg map[string]string) {
	context.JSON(412, gin.H{
		"error":     1,
		"message":   msg,
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
}
