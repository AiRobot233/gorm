package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
	"gorm.io/gorm"
)

// UserList 列表
func UserList(page string, pageSize string, params model.UserSearch) (bool, any) {
	var users []model.User //定义表结构
	var count int64
	db.Table("user").Scopes(model.UserSearchFunc(params)).Count(&count)
	result := db.Table("user").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Select([]string{"id", "name", "phone", "role_id", "status", "created_at"}).Scopes(model.UserSearchFunc(params)).Scopes(model.Paginate(page, pageSize)).Find(&users)
	return utils.R(result, utils.P(users, count))
}

// UserAdd 新增
func UserAdd(params validate.User) (bool, any) {
	//判断值是否存在
	if params.Password == "" {
		params.Password = "Aa@112233"
	}
	salt := utils.GetSalt(params.Password)
	user := model.User{
		Name:     params.Name,
		Phone:    params.Phone,
		Salt:     salt,
		Password: utils.Md5(params.Password + salt),
		RoleId:   params.RoleId,
		Status:   params.Status,
	}
	result := db.Create(&user)
	return utils.R(result, nil)
}

// UserEdit 修改
func UserEdit(id string, params validate.User) (bool, any) {
	user := model.User{}
	res := db.First(&user, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	user.Name = params.Name
	user.Phone = params.Phone
	user.RoleId = params.RoleId
	user.Status = params.Status
	if params.Password != "" {
		bol, res := utils.SetPwd(params.Password, user.Salt)
		if bol != true {
			return false, res
		}
		user.Password = res
	}
	result := db.Save(&user)
	return utils.R(result, nil)
}

// UserDel 删除
func UserDel(id string) (bool, any) {
	user := model.User{}
	res := db.First(&user, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	result := db.Delete(&user)
	return utils.R(result, nil)
}
