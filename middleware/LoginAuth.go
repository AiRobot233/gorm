package middleware

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func LoginAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("Authorization")
		if token == "" {
			utils.Error(context, "未登录", 401)
			context.Abort()
			return
		} else {
			err, data := utils.CheckJwt(token)
			if err != nil {
				utils.Error(context, data, 401)
				context.Abort()
				return
			} else {
				//数据 进入协程上下文
				context.Set("user", data)
			}
		}
		context.Next()
	}
}
