package main

import (
	"gin/controller"
	"gin/middleware"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/storage", "./storage") //文件访问配置地址
	v1 := router.Group("/v1")
	{
		v1.POST("/someGet", func(context *gin.Context) {
			bol := validate.LoginValidate(context)
			if bol {
				controller.Login(context)
			}
		})
		v1.POST("/redis", func(context *gin.Context) {
			controller.Redis()
		})
		v1.GET("/jwt", func(context *gin.Context) {
			controller.Jwt(context)
		})
		//新增数据
		v1.POST("/save", func(context *gin.Context) {
			controller.SqlSave(context)
		})
		//查询数据
		v1.GET("/get", func(context *gin.Context) {
			controller.SqlGet(context)
		})
		//删除数据
		v1.DELETE("/del/:id", func(context *gin.Context) {
			controller.SqlDel(context)
		})
		//测试
		v1.GET("/index", func(context *gin.Context) {
			controller.Index()
		})
		//上传文件
		v1.POST("/upload", func(context *gin.Context) {
			controller.Upload(context)
		})
	}

	v2 := router.Group("/v2")
	{
		v2.Use(middleware.LoginAuth()).POST("/check_jwt", func(context *gin.Context) {
			controller.A(context)
		})
	}
	_ = router.Run(":8080")
}
