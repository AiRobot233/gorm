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
