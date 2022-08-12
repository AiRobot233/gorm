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
	_, token := utils.GetJwt(data, 21)
	utils.Success(context, token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, data)
}

func Test(c *gin.Context) {
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.ShouldBind(&params); err == nil {
		utils.Success(c, params)
	} else {
		utils.Error(c, err.Error())
	}
}
