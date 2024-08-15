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

// RoleTree 获取角色，递归输出树状
func RoleTree(menus []*model.RoleTree, pid int) []*model.RoleTree {
	//定义子节点目录
	var nodes []*model.RoleTree
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for _, v := range menus {
			//查询所有该菜单下的所有子菜单
			if v.Pid == pid {
				//特别注意压入元素不是单个所有加三个 **...** 告诉切片无论多少元素一并压入
				v.Child = append(v.Child, RoleTree(menus, v.Id)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}

// DictionaryTree 获取字典，递归输出树状
func DictionaryTree(menus []*model.DictionaryTree, pid int) []*model.DictionaryTree {
	//定义子节点目录
	var nodes []*model.DictionaryTree
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for _, v := range menus {
			//查询所有该菜单下的所有子菜单
			if v.Pid == pid {
				//特别注意压入元素不是单个所有加三个 **...** 告诉切片无论多少元素一并压入
				v.Child = append(v.Child, DictionaryTree(menus, v.Id)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}

// UnitTree 单位，递归输出树状
func UnitTree(menus []*model.UnitTree, pid int) []*model.UnitTree {
	//定义子节点目录
	var nodes []*model.UnitTree
	if reflect.ValueOf(menus).IsValid() {
		//循环所有一级菜单
		for _, v := range menus {
			//查询所有该菜单下的所有子菜单
			if v.Pid == pid {
				//特别注意压入元素不是单个所有加三个 **...** 告诉切片无论多少元素一并压入
				v.Child = append(v.Child, UnitTree(menus, v.Id)...)
				nodes = append(nodes, v)
			}
		}
	}
	return nodes
}
