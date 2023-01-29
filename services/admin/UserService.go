package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
	"gorm.io/gorm"
)

// UserList 列表
func UserList(page string, pageSize string) (bool, interface{}) {
	var users []model.User //定义表结构
	var count int64
	db.Table("user").Count(&count)
	result := db.Debug().Table("user").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Select([]string{"id", "name", "phone", "role_id", "status", "created_at"}).Scopes(model.Paginate(page, pageSize)).Find(&users)
	return utils.R(result, utils.P(users, count))
}

// UserAdd 新增
func UserAdd(params validate.User) (bool, interface{}) {
	//判断值是否存在
	var status int
	if params.Password == "" {
		params.Password = "Aa@112233"
	}
	if params.Status == 0 {
		params.Status = 1
	}
	salt := utils.GetSalt(params.Password)
	user := model.User{
		Name:     params.Name,
		Phone:    params.Phone,
		Salt:     salt,
		Password: utils.Md5(params.Password + salt),
		RoleId:   params.RoleId,
		Status:   status,
	}
	result := db.Create(&user)
	return utils.R(result, nil)
}

// UserEdit 修改
func UserEdit(id string, params validate.User) (bool, interface{}) {
	user := model.User{}
	db.First(&user, id)
	user.Name = params.Name
	user.Phone = params.Phone
	user.RoleId = params.RoleId
	user.Status = params.Status
	if params.Password != "" {
		bol, res := SetPwd(params.Password, user.Salt)
		if bol != true {
			return false, res
		}
		user.Password = res
	}
	result := db.Save(&user)
	return utils.R(result, nil)
}

// UserDel 删除
func UserDel(id string) (bool, interface{}) {
	user := model.User{}
	result := db.Delete(&user, id)
	return utils.R(result, nil)
}

// SetPwd 修改密码操作
func SetPwd(password string, salt string) (bool, string) {
	err := utils.CheckPasswordLever(password) //校验密码安全性
	if err != nil {
		return false, err.Error()
	}
	return true, utils.Md5(password + salt)
}
