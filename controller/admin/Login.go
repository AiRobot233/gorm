package admin

import (
	"gin/services/admin"
	"gin/utils"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

// Login 登录返回token
func Login(c *gin.Context) {
	params := validate.Login{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.Login(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}

}
