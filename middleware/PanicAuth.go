package middleware

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func PanicAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() { // 必须要先声明defer，否则不能捕获到panic异常
			if err := recover(); err != nil {
				utils.Error(context, err)
				context.Abort()
				return
			}
		}()
		context.Next()
	}
}
