package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// RoleList 角色列表
func RoleList() (bool, any) {
	var role []*model.RoleTree
	result := db.Find(&role)
	return utils.R(result, RoleTree(role, 0))
}

// RoleAdd 角色新增
func RoleAdd(params validate.Role) (bool, any) {
	role := model.Role{}
	role.RoleSetFromData(params)
	result := db.Create(&role)
	return utils.R(result, nil)
}

// RoleEdit 角色修改
func RoleEdit(id string, params validate.Role) (bool, any) {
	role := model.Role{}
	res := db.First(&role, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	role.RoleSetFromData(params)
	result := db.Save(&role)
	return utils.R(result, nil)
}

// RoleDel 角色删除
func RoleDel(id string) (bool, any) {
	role := model.Role{}
	res := db.First(&role, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	result := db.Delete(&role)
	return utils.R(result, nil)
}

// RoleSelect 下拉
func RoleSelect() (bool, any) {
	var role []*model.RoleTree
	result := db.Find(&role)
	return utils.R(result, RoleTree(role, 0))
}
