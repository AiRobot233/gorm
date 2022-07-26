package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

//角色列表
func RoleList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "2")
	bol, res := admin.RoleList(page, pageSize)
	utils.Send(c, bol, res)
}

//角色添加
func RoleAdd(c *gin.Context) {
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.Bind(&params); err == nil {
		bol, data := admin.RoleAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

//角色修改
func RoleEdit(c *gin.Context) {
	id := c.Param("id")
	var params map[string]interface{}     //声明变量，不分配内存
	params = make(map[string]interface{}) //必可不少，分配内存
	if err := c.BindJSON(&params); err == nil {
		bol, res := admin.RoleEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

//角色删除
func RoleDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.RoleDel(id)
	utils.Send(c, bol, res)
}
