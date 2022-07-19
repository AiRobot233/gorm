package admin

import (
	"gin/model"
	"gin/utils"
	"reflect"
)

//规则列表 树状
func RuleList() (bool, interface{}) {
	var rule []*model.RuleTree
	result := db.Find(&rule)
	return utils.R(result, ruleTree(rule, 0))
}

//规则添加
func RuleAdd(params map[string]interface{}) (bool, interface{}) {
	rule := model.Rule{}
	model.RuleSetFromData(&rule, params)
	result := db.Create(&rule)
	return utils.R(result, nil)
}

//规则修改
func RuleEdit(id string, params map[string]interface{}) (bool, interface{}) {
	rule := model.Rule{}
	db.First(&rule, id)
	model.RuleSetFromData(&rule, params)
	result := db.Save(&rule)
	return utils.R(result, nil)
}

//规则删除
func RuleDel(id string) (bool, interface{}) {
	rule := model.Rule{}
	result := db.Delete(&rule, id)
	return utils.R(result, nil)
}

//递归输出树状(内部包方法)
func ruleTree(menus []*model.RuleTree, pid int) []*model.RuleTree {
	//定义子节点目录
	var nodes []*model.RuleTree
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for _, v := range menus {
			//查询所有该菜单下的所有子菜单
			if v.Pid == pid {
				//特别注意压入元素不是单个所有加三个 **...** 告诉切片无论多少元素一并压入
				v.Child = append(v.Child, ruleTree(menus, v.Id)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}
