package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func RuleList(c *gin.Context) {
	bol, res := admin.RuleList()
	utils.Send(c, bol, res)
}

//规则添加
func RuleAdd(c *gin.Context) {
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.RuleAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

//规则修改
func RuleEdit(c *gin.Context) {
	id := c.Param("id")
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.RuleEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

//规则删除
func RuleDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.RuleDel(id)
	utils.Send(c, bol, res)
}

//规则下拉
func RuleSelect(c *gin.Context) {
	types := c.DefaultQuery("type", "")
	bol, res := admin.RuleSelect(types)
	utils.Send(c, bol, res)
}
