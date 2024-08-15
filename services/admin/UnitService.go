package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

func UnitList() (bool, any) {
	var params []*model.UnitTree
	result := db.Order("sort asc,id asc").Find(&params)
	return utils.R(result, UnitTree(params, 0))
}

// UnitSelect 组件下拉
func UnitSelect(isRegister bool) (bool, any) {
	var tree []*model.UnitTree
	if isRegister {
		result := db.Order("sort asc,id asc").Where("is_register = 1").Find(&tree)
		return utils.R(result, UnitTree(tree, 0))
	} else {
		result := db.Order("sort asc,id asc").Find(&tree)
		return utils.R(result, UnitTree(tree, 0))
	}
}

// UnitAdd 添加
func UnitAdd(params validate.Unit) (bool, any) {
	res := model.Unit{}
	res.UnitSetFromData(params)
	result := db.Create(&res)
	return utils.R(result, nil)
}

// UnitEdit 修改
func UnitEdit(id string, params validate.Unit) (bool, any) {
	unit := model.Unit{}
	res := db.First(&unit, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	unit.UnitSetFromData(params)
	result := db.Save(&unit)
	return utils.R(result, nil)
}

// UnitDel 规则删除
func UnitDel(id string) (bool, any) {
	rule := model.Unit{}
	res := db.First(&rule, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	result := db.Delete(&rule)
	return utils.R(result, nil)
}

func UnitListApp() (bool, any) {
	var content []model.Unit //定义表结构
	result := db.Table("unit").Where("is_unit = 1").Find(&content)
	return utils.R(result, content)
}
