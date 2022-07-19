package admin

import (
	"gin/model"
	"gin/utils"
)

//角色列表
func RoleList(page string, pageSize string) (bool, interface{}) {
	role := model.Role{}
	var count int64
	result := db.Table("role").Scopes(model.Paginate(page, pageSize)).Scan(&role).Count(&count)
	return utils.R(result, utils.P(role, count))
}

//角色新增
func RoleAdd(params map[string]interface{}) (bool, interface{}) {
	role := model.Role{}
	model.RoleSetFromData(&role, params)
	result := db.Create(&role)
	return utils.R(result, nil)
}

//角色修改
func RoleEdit(id string, params map[string]interface{}) (bool, interface{}) {
	role := model.Role{}
	db.First(&role, id)
	model.RoleSetFromData(&role, params)
	result := db.Save(&role)
	return utils.R(result, nil)
}

//角色删除
func RoleDel(id string) (bool, interface{}) {
	role := model.Role{}
	res := db.Where("id = ?", id).First(&role)
	if res.RowsAffected == 0 {
		return false, "数据不存在"
	}
	result := db.Delete(&role)
	return utils.R(result, nil)
}
