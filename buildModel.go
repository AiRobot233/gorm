package main

import (
	"github.com/pulingfu/tblschema"
)

func main() {
	//简单用法
	simple := tblschema.NewTblToStructHandler()
	simple.
		SetDsn("root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8").
		SetTableName("user").
		//默认路径为当前目录
		SetSavePath("model/User.go").GenerateTblStruct()
	// SetPackageInfo("plf_test_package", "", "").s
}
