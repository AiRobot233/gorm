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
	var data map[string]interface{}     //定义map
	data = make(map[string]interface{}) //初始化map
	data["routes"] = RuleTree(rules, 0)
	data["roles"] = roles
	return utils.R(result, data)
}
