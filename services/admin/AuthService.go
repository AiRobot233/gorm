package admin

import (
	"gin/model"
	"gin/utils"
	"gorm.io/gorm"
	"strings"
)

type Roles struct {
	Router    string `json:"router"`
	Operation string `json:"operation"`
}

//获取登录人员权限
func GetRoutes(user map[string]interface{}) (bool, interface{}) {
	role := model.Role{}
	res := db.First(&role, user["role_id"]) //查询角色
	if res.RowsAffected == 0 {
		return false, "角色不存在"
	}
	var rules []*model.RuleTree
	var result *gorm.DB

	var roles []Roles
	if role.IsSystem == 2 {
		result = db.Where("type = ?", "page").Order("sort desc").Find(&rules) //查询所有规则
		db.Raw("SELECT b.router,a.operation FROM rule AS b LEFT JOIN (SELECT pid,GROUP_CONCAT(method SEPARATOR ',') AS operation FROM `rule` WHERE type = 'api' GROUP BY pid) AS a ON a.pid = b.id WHERE a.operation IS NOT NULL").Find(&roles)
	} else {
		result = db.Where("id IN ? AND type = ?", strings.Split(role.Rule, `,`), "page").Order("sort desc").Find(&rules) //查询规则
		db.Raw("SELECT b.router,a.operation FROM rule AS b LEFT JOIN (SELECT pid,GROUP_CONCAT(method SEPARATOR ',') AS operation FROM `rule` WHERE type = 'api' AND id IN ? GROUP BY pid) AS a ON a.pid = b.id WHERE a.operation IS NOT NULL", strings.Split(role.Rule, `,`)).Find(&roles)
	}
	data := utils.GetSlice()
	data["routes"] = RuleTree(rules, 0)
	data["roles"] = roles
	return utils.R(result, data)
}

//修改自己密码
func ChangePwd(params map[string]interface{}, user map[string]interface{}) (bool, interface{}) {
	oldPassword := params["old_password"].(string)
	password := params["password"].(string)
	u := model.User{}
	result := db.First(&u, user["id"].(float64))
	if result.RowsAffected <= 0 {
		return false, "用户未找到！"
	}
	if utils.Md5(oldPassword+u.Salt) == u.Password {
		//旧密码正确
		bol, data := SetPwd(password, u.Salt)
		if bol != true {
			return false, data
		}
		u.Password = data
		res := db.Save(&u)
		return utils.R(res, nil)
	} else {
		return false, "旧密码错误！"
	}
}
