package controller

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Jwt(context *gin.Context) {
	data := utils.GetSlice()
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
	params := utils.GetSlice()
	if err := c.ShouldBind(&params); err == nil {
		utils.Success(c, params)
	} else {
		utils.Error(c, err.Error())
	}
}
