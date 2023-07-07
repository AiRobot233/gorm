package middleware

import (
	"gin/model"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var db = model.GetDb()

func RuleAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取前端访问的路由信息
		url := c.FullPath()
		method := c.Request.Method
		//查询用户是否有权限访问接口
		user, err := c.Get("user")
		if err {
			var u *model.User
			data := user.(map[string]any)
			db.First(&u, data["id"])
			if u.Id <= 0 {
				utils.Error(c, "用户已被删除", 401)
				c.Abort()
				return
			}
			//查询规则
			var rule *model.Rule
			db.Where("router = ? AND method = ?", url, method).First(&rule)
			if rule.Id <= 0 {
				utils.Error(c, "接口地址未配置")
				c.Abort()
				return
			}
			//查询角色
			var role *model.Role
			db.Where("id = ?", u.RoleId).First(&role)
			if role.Id <= 0 {
				utils.Error(c, "角色已被删除，请联系管理员修复")
				c.Abort()
				return
			}
			if role.IsSystem == 1 {
				res := utils.InArray(strings.Split(role.Rule, `,`), strconv.Itoa(rule.Id))
				if !res {
					utils.Error(c, "暂无权限")
					c.Abort()
					return
				}
			}
		} else {
			utils.Error(c, "用户身份丢失！请重新登录", 401)
			c.Abort()
			return
		}
		c.Next()
	}
}
