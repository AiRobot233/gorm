package controller

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Jwt(context *gin.Context) {
	var data map[string]interface{}     //定义map
	data = make(map[string]interface{}) //初始化map
	data["uid"] = 1
	data["user_name"] = "hhh"
	token := utils.GetJwt(data, 21)
	utils.Success(context, token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, data)
}
