package model

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/pulingfu/tblschema"
)

func Build(c *gin.Context) {
	params := utils.GetSlice()
	if err := c.ShouldBindJSON(&params); err == nil {
		info(params)
		utils.Send(c, true, nil)
	} else {
		utils.Error(c, err.Error())
	}
}

func info(params map[string]any) {
	name := params["table_name"].(string)
	path := params["path"].(string)
	//简单用法
	simple := tblschema.NewTblToStructHandler()
	simple.
		SetDsn("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8").
		SetTableName(name).
		//默认路径为当前目录
		SetSavePath(path).GenerateTblStruct()
	// SetPackageInfo("plf_test_package", "", "").s
}
