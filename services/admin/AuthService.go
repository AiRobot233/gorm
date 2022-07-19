package admin

import (
	"gin/model"
	"gin/utils"
	"gorm.io/gorm"
)

//获取登录人员权限
func GetAuth(user map[string]interface{}) (bool, interface{}) {
	role := model.Role{}
	res := db.First(&role, user["role_id"]) //查询角色
	if res.RowsAffected == 0 {
		return false, "角色不存在"
	}
	var rules []*model.RuleTree
	var result *gorm.DB
	if role.Rule == "*" {
		result = db.Find(&rules) //查询所有规则
	} else {
		result = db.Where("id in ?", role.Rule).Find(&rules) //查询规则
	}
	return utils.R(result, RuleTree(rules, 0))
}
