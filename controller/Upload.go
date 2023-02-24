package controller

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"os"
	"path"
	"strconv"
	"time"
)

func Upload(context *gin.Context) {
	// 单文件
	file, _ := context.FormFile("file")
	//获取文件的后缀名
	ext := path.Ext(file.Filename)
	//获取当前时间
	fileNameInt := time.Now().UnixNano()
	rand := strconv.FormatInt(fileNameInt, 10) //转str
	// 上传文件到指定的路径
	dst := build() + "/" + rand + ext
	err := context.SaveUploadedFile(file, dst)
	if err != nil {
		utils.Error(context, err.Error())
	} else {
		utils.Success(context, "")
	}
}

// 创建文件夹
func build() string {
	day := time.Now().Format(time.DateOnly)
	filePath := "./storage/" + day
	//判断文件夹是否存在
	_, err := os.Stat(filePath)
	if err != nil {
		_ = os.Mkdir(filePath, os.ModePerm)
	}
	return filePath
}
