package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

//登录返回token
func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	bol, data := admin.Login(phone, password)
	utils.Send(c, bol, data)
}
