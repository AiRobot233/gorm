package admin

import (
	"gin/services/admin"
	"gin/utils"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

func UnitList(c *gin.Context) {
	bol, res := admin.UnitList()
	utils.Send(c, bol, res)
}

// UnitAdd 添加
func UnitAdd(c *gin.Context) {
	params := validate.Unit{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.UnitAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

// UnitEdit 修改
func UnitEdit(c *gin.Context) {
	id := c.Param("id")
	params := validate.Unit{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.UnitEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

// UnitDel 删除
func UnitDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.UnitDel(id)
	utils.Send(c, bol, res)
}
