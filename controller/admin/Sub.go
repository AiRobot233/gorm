package admin

import (
	"gin/services/admin"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Assembly(c *gin.Context) {
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		bol, data := admin.Assembly(params)
		utils.Send(c, bol, data)
	} else {
		utils.Error(c, err.Error())
	}
}
