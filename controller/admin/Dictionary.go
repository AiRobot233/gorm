package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

//字典列表
func DictionaryList(c *gin.Context) {
	bol, data := admin.DictionaryList()
	utils.Send(c, bol, data)
}

//字典添加
func DictionaryAdd(c *gin.Context) {
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.DictionaryAdd(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}

//字典修改
func DictionaryEdit(c *gin.Context) {
	id := c.Param("id")
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, res := admin.DictionaryEdit(id, params)
		utils.Send(c, bol, res)
	} else {
		utils.Error(c, err.Error())
	}
}

//字典删除
func DictionaryDel(c *gin.Context) {
	id := c.Param("id")
	bol, res := admin.DictionaryDel(id)
	utils.Send(c, bol, res)
}

//获取字典数据（不鉴权）
func UnitDictionary(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	bol, res := admin.UnitDictionary(name)
	utils.Send(c, bol, res)
}

//字典下拉
func DictionarySelect(c *gin.Context) {
	bol, res := admin.DictionarySelect()
	utils.Send(c, bol, res)
}
