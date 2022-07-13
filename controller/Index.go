package controller

import (
	"fmt"
	"gin/model"
	"gin/utils"
	"github.com/gin-gonic/gin"
)

func Jwt(context *gin.Context) {
	var data map[string]interface{}     //定义map
	data = make(map[string]interface{}) //初始化map
	data["uid"] = 1
	data["user_name"] = "hhh"
	token := utils.GetJwt(data, 21)
	utils.Success(context, token)
}

func A(context *gin.Context) {
	data, _ := context.Get("user")
	utils.Success(context, data)
}

func SqlSave(context *gin.Context) {
	user := model.User{}
	user.Name = "张三"
	user.Age = 25
	result := user.Db().Create(&user)
	if result.Error == nil {
		utils.Success(context, user.Id)
	} else {
		utils.Error(context, result.Error)
	}
}

func SqlGet(context *gin.Context) {
	user := model.User{}
	result := user.Db().Preload("Books", "id != (?)", 1).First(&user)
	if result.Error == nil {
		utils.Success(context, user)
	} else {
		utils.Error(context, result.Error)
	}
}

func SqlDel(context *gin.Context) {
	id := context.Param("id")
	utils.Success(context, id)
}

func Index() {
	var a = 1
	fmt.Println("a=", a)
	fmt.Println("&a=", &a)
}
