package utils

import (
	"github.com/gin-gonic/gin"
	"time"
)

var successMsg = "success" //成功默认值
var errorCode = 200        //失败code码

func Success(c *gin.Context, data any, msg ...string) {
	if len(msg) > 0 {
		successMsg = msg[0]
	}
	c.JSON(200, gin.H{
		"error":     0,
		"message":   successMsg,
		"data":      data,
		"timestamp": time.Now().Unix(),
	})
}

func Error(c *gin.Context, msg any, code ...int) {
	if len(code) > 0 {
		errorCode = code[0]
	}
	c.JSON(errorCode, gin.H{
		"error":     1,
		"message":   msg,
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
}

// ValidateError 表单验证错误返回
func ValidateError(c *gin.Context, msg map[string]string) {
	c.JSON(412, gin.H{
		"error":     1,
		"message":   msg,
		"data":      nil,
		"timestamp": time.Now().Unix(),
	})
}

// Send 封装输出数据
func Send(c *gin.Context, bol bool, data any) {
	if bol {
		Success(c, data)
	} else {
		Error(c, data)
	}
}
