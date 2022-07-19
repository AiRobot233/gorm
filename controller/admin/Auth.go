package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	user, err := c.Get("user")
	if err {
		bol, data := admin.GetAuth(user.(map[string]interface{}))
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}
