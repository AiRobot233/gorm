package admin

import (
	"gin/services/admin"
	"gin/utils"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

// UserList 列表
func UserList(c *gin.Context) {
	page := c.DefaultQuery("page", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	bol, res := admin.UserList(page, pageSize)
	utils.Send(c, bol, res)
}

// UserAdd 新增
func UserAdd(c *gin.Context) {
	params := validate.User{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.UserAdd(params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

// UserEdit 修改
func UserEdit(c *gin.Context) {
	id := c.Param("id")
	params := validate.User{}
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.UserEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

// UserDel 删除
func UserDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.UserDel(id)
	utils.Send(c, bol, res)
}
