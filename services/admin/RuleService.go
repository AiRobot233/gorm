package admin

import (
	"gin/model"
	"gin/utils"
)

// RuleList 规则列表 树状
func RuleList() (bool, interface{}) {
	var rule []*model.RuleTree
	result := db.Find(&rule)
	return utils.R(result, RuleTree(rule, 0))
}

// RuleAdd 规则添加
func RuleAdd(params map[string]interface{}) (bool, interface{}) {
	rule := model.Rule{}
	rule.RuleSetFromData(params)
	result := db.Create(&rule)
	return utils.R(result, nil)
}

// RuleEdit 规则修改
func RuleEdit(id string, params map[string]interface{}) (bool, interface{}) {
	rule := model.Rule{}
	db.First(&rule, id)
	rule.RuleSetFromData(params)
	result := db.Save(&rule)
	return utils.R(result, nil)
}

// RuleDel 规则删除
func RuleDel(id string) (bool, interface{}) {
	rule := model.Rule{}
	res := db.Where("id = ?", id).First(&rule)
	if res.RowsAffected == 0 {
		return false, "数据不存在"
	}
	result := db.Delete(&rule)
	return utils.R(result, nil)
}

func RuleSelect(types string) (bool, interface{}) {
	var rule []*model.RuleTree
	var where model.Rule
	if types != "" {
		where.Type = types
	}
	result := db.Find(&rule, where)
	return utils.R(result, RuleTree(rule, 0))
}
