package admin

import (
	"gin/model"
	"reflect"
)

var db = model.GetDb() //定义db链接

// RuleTree 获取规则，递归输出树状
func RuleTree(menus []*model.RuleTree, pid int) []*model.RuleTree {
	//定义子节点目录
	var nodes []*model.RuleTree
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for _, v := range menus {
			//查询所有该菜单下的所有子菜单
			if v.Pid == pid {
				//特别注意压入元素不是单个所有加三个 **...** 告诉切片无论多少元素一并压入
				v.Child = append(v.Child, RuleTree(menus, v.Id)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}
