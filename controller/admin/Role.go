package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

// RoleList 角色列表
func RoleList(c *gin.Context) {
	bol, res := admin.RoleList()
	utils.Send(c, bol, res)
}

// RoleAdd 角色添加
func RoleAdd(c *gin.Context) {
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.RoleAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

// RoleEdit 角色修改
func RoleEdit(c *gin.Context) {
	id := c.Param("id")
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.RoleEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

// RoleDel 角色删除
func RoleDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.RoleDel(id)
	utils.Send(c, bol, res)
}

// RoleSelect 下拉
func RoleSelect(c *gin.Context) {
	bol, res := admin.RoleSelect()
	utils.Send(c, bol, res)
}
