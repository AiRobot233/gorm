package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// RuleList 规则列表 树状
func RuleList() (bool, any) {
	var rule []*model.RuleTree
	result := db.Find(&rule)
	return utils.R(result, RuleTree(rule, 0))
}

// RuleAdd 规则添加
func RuleAdd(params validate.Rule) (bool, any) {
	rule := model.Rule{}
	rule.RuleSetFromData(params)
	result := db.Create(&rule)
	return utils.R(result, nil)
}

// RuleEdit 规则修改
func RuleEdit(id string, params validate.Rule) (bool, any) {
	rule := model.Rule{}
	res := db.First(&rule, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	rule.RuleSetFromData(params)
	result := db.Save(&rule)
	return utils.R(result, nil)
}

// RuleDel 规则删除
func RuleDel(id string) (bool, any) {
	rule := model.Rule{}
	res := db.First(&rule, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	result := db.Delete(&rule)
	return utils.R(result, nil)
}

func RuleSelect(types string) (bool, any) {
	var rule []*model.RuleTree
	var where model.Rule
	if types != "" {
		where.Type = types
	}
	result := db.Find(&rule, where)
	return utils.R(result, RuleTree(rule, 0))
}
