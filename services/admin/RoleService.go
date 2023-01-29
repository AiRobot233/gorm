package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// RoleList 角色列表
func RoleList() (bool, interface{}) {
	var role []*model.RoleTree
	result := db.Find(&role)
	return utils.R(result, RoleTree(role, 0))
}

// RoleAdd 角色新增
func RoleAdd(params validate.Role) (bool, interface{}) {
	role := model.Role{}
	role.RoleSetFromData(params)
	result := db.Create(&role)
	return utils.R(result, nil)
}

// RoleEdit 角色修改
func RoleEdit(id string, params validate.Role) (bool, interface{}) {
	role := model.Role{}
	db.First(&role, id)
	role.RoleSetFromData(params)
	result := db.Save(&role)
	return utils.R(result, nil)
}

// RoleDel 角色删除
func RoleDel(id string) (bool, interface{}) {
	role := model.Role{}
	res := db.Where("id = ?", id).First(&role)
	if res.RowsAffected == 0 {
		return false, "数据不存在"
	}
	result := db.Delete(&role)
	return utils.R(result, nil)
}

// RoleSelect 下拉
func RoleSelect() (bool, interface{}) {
	var role []*model.RoleTree
	result := db.Find(&role)
	return utils.R(result, RoleTree(role, 0))
}
