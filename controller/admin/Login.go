package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

//登录返回token
func Login(c *gin.Context) {
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.Login(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}

}
