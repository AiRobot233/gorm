package admin

import (
	"gin/model"
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func SettingWeb(c *gin.Context) {
	bol, data := admin.SettingWeb()
	utils.Send(c, bol, data)
}

func SettingList(c *gin.Context) {
	bol, data := admin.SettingList()
	utils.Send(c, bol, data)
}

func SettingSave(c *gin.Context) {
	var query []model.Setting
	// 绑定 JSON 数组到切片
	if err := c.ShouldBindJSON(&query); err != nil {
		utils.Error(c, "参数获取失败")
		return
	}
	bol, data := admin.SettingSave(query)
	utils.Send(c, bol, data)
}
