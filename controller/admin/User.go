package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

//列表
func UserList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "2")
	bol, res := admin.UserList(page, pageSize)
	utils.Send(c, bol, res)
}

//新增
func UserAdd(c *gin.Context) {
	var params map[string]interface{}
	params = make(map[string]interface{})
	if err := c.BindJSON(&params); err == nil {
		bol, res := admin.UserAdd(params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

//修改
func UserEdit(c *gin.Context) {
	id := c.Param("id")
	var params map[string]interface{}
	params = make(map[string]interface{})
	if err := c.BindJSON(&params); err == nil {
		bol, res := admin.UserEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

//删除
func UserDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.UserDel(id)
	utils.Send(c, bol, res)
}
