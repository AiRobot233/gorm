package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func GetRoutes(c *gin.Context) {
	user, err := c.Get("user")
	if err {
		bol, data := admin.GetRoutes(user.(map[string]interface{}))
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}

//修改自己密码
func ChangePwd(c *gin.Context) {
	user, err := c.Get("user")
	if err {
		var params map[string]interface{}     //声明变量，不分配内存
		params = make(map[string]interface{}) //必可不少，分配内存
		if err := c.ShouldBindJSON(&params); err == nil {
			bol, data := admin.ChangePwd(params, user.(map[string]interface{}))
			utils.Send(c, bol, data)
		} else {
			utils.Error(c, err.Error())
		}

	} else {
		utils.Error(c, "用户身份丢失！请重新登录", 401)
	}
}
