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
		url := strings.Split(c.Request.URL.String(), "?")[0]
		method := c.Request.Method
		//判断PUT DELETE 删除restful的id参数
		if method == "PUT" || method == "DELETE" {
			newUrl := ""
			arr := strings.Split(url, "/")
			for i := 0; i < len(arr)-1; i++ {
				if arr[i] != "" {
					newUrl = newUrl + "/" + arr[i]
				}
			}
			url = newUrl //重新赋值
		}
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
