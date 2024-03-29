package admin

import (
	"gin/services/admin"
	"gin/utils"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

func RuleList(c *gin.Context) {
	bol, res := admin.RuleList()
	utils.Send(c, bol, res)
}

// RuleAdd 规则添加
func RuleAdd(c *gin.Context) {
	params := validate.Rule{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.RuleAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

// RuleEdit 规则修改
func RuleEdit(c *gin.Context) {
	id := c.Param("id")
	params := validate.Rule{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.RuleEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

// RuleDel 规则删除
func RuleDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.RuleDel(id)
	utils.Send(c, bol, res)
}
