package controller

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Upload(context *gin.Context) {
	// 单文件
	file, _ := context.FormFile("file")
	// 上传文件到指定的路径
	dst := "./storage/" + file.Filename
	err := context.SaveUploadedFile(file, dst)
	if err != nil {
		utils.Error(context, err)
	} else {
		utils.Success(context, "")
	}
}
