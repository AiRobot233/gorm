package main

import (
	"gin/controller"
	"gin/controller/admin"
	"gin/middleware"
	"gin/validate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/storage", "./storage") //文件访问配置地址
	a := router.Group("/admin")
	{
		a.POST("/login", func(c *gin.Context) {
			bol := validate.LoginValidate(c)
			if bol {
				admin.Login(c)
			}
		})

		a.POST("/test", func(c *gin.Context) {
			controller.Test(c)
		})

		notLogin := a.Use(middleware.LoginAuth())
		//角色下拉
		notLogin.GET("/role/select", func(c *gin.Context) {
			admin.RoleSelect(c)
		})
		//规则下拉
		notLogin.GET("/rule/select", func(c *gin.Context) {
			admin.RuleSelect(c)
		})

		//鉴权
		auth := a.Use(middleware.LoginAuth())
		//获取用户权限
		auth.GET("/routes", func(c *gin.Context) {
			admin.GetRoutes(c)
		})

		//用户列表
		auth.GET("/user", func(c *gin.Context) {
			admin.UserList(c)
		})
		//用户添加
		auth.POST("/user", func(c *gin.Context) {
			bol := validate.UserValidate(c)
			if bol {
				admin.UserAdd(c)
			}
		})
		//用户修改
		auth.PUT("/user/:id", func(c *gin.Context) {
			bol := validate.UserValidate(c)
			if bol {
				admin.UserEdit(c)
			}
		})
		//用户删除
		auth.DELETE("/user/:id", func(c *gin.Context) {
			admin.UserDel(c)
		})

		//规则列表
		auth.GET("/rule", func(c *gin.Context) {
			admin.RuleList(c)
		})
		//规则添加
		auth.POST("/rule", func(c *gin.Context) {
			bol := validate.RuleValidate(c)
			if bol {
				admin.RuleAdd(c)
			}
		})
		//规则修改
		auth.PUT("/rule/:id", func(c *gin.Context) {
			bol := validate.RuleValidate(c)
			if bol {
				admin.RuleEdit(c)
			}
		})
		//规则删除
		auth.DELETE("/rule/:id", func(c *gin.Context) {
			admin.RuleDel(c)
		})

		//角色列表
		auth.GET("/role", func(c *gin.Context) {
			admin.RoleList(c)
		})
		//角色新增
		auth.POST("/role", func(c *gin.Context) {
			bol := validate.RoleValidate(c)
			if bol {
				admin.RoleAdd(c)
			}
		})
		//角色修改
		auth.PUT("/role/:id", func(c *gin.Context) {
			bol := validate.RoleValidate(c)
			if bol {
				admin.RoleEdit(c)
			}
		})
		//角色删除
		auth.DELETE("/role/:id", func(c *gin.Context) {
			admin.RoleDel(c)
		})

		//字典列表
		auth.GET("/dictionary", func(c *gin.Context) {
			admin.DictionaryList(c)
		})
		//字典新增
		auth.POST("/dictionary", func(c *gin.Context) {

		})
		//字典修改
		auth.PUT("/dictionary/:id", func(c *gin.Context) {

		})
		//字典删除
		auth.DELETE("/dictionary/:id", func(c *gin.Context) {

		})

		//上传文件
		auth.POST("/upload", func(c *gin.Context) {
			controller.Upload(c)
		})
	}
	_ = router.Run(":9501")
}
