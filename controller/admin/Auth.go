package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func GetRoutes(c *gin.Context) {
	user, err := c.Get("user")
	if err {
		bol, data := admin.GetRoutes(user.(map[string]any))
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}

// ChangePwd 修改自己密码
func ChangePwd(c *gin.Context) {
	user, err := c.Get("user")
	if err {
		params := utils.GetSlice()
		if err := c.ShouldBindJSON(&params); err == nil {
			bol, data := admin.ChangePwd(params, user.(map[string]any))
			utils.Send(c, bol, data)
		} else {
			utils.Error(c, err.Error())
		}
	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}

// FirstPwd 第一次修改密码
func FirstPwd(c *gin.Context) {
	userAny, err := c.Get("user")
	if err {
		var params struct {
			Password string `form:"password"`
		}
		if err := c.ShouldBind(&params); err == nil {
			user := userAny.(map[string]any)
			userId := user["id"].(float64)
			bol, data := admin.FirstPwd(userId, params.Password)
			utils.Send(c, bol, data)
		} else {
			utils.Error(c, err.Error())
		}
	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}
