package admin

import (
	"gin/model"
	"gin/utils"
	"gorm.io/gorm"
)

// UserList 列表
func UserList(page string, pageSize string, user map[string]interface{}) (bool, interface{}) {
	var users []model.User //定义表结构
	var count int64
	db.Table("user").Count(&count)
	result := db.Debug().Table("user").Preload("Role", func(db *gorm.DB) *gorm.DB {
		return db.Select([]string{"id", "name"})
	}).Select([]string{"id", "name", "phone", "role_id", "status", "created_at"}).Scopes(model.Paginate(page, pageSize)).Find(&users)
	return utils.R(result, utils.P(users, count))
}

// UserAdd 新增
func UserAdd(params map[string]interface{}) (bool, interface{}) {
	//判断值是否存在
	var status int
	if _, ok := params["password"]; !ok {
		params["password"] = "Aa@112233"
	}
	if _, ok := params["status"]; !ok {
		status = 1
	} else {
		status = int(params["status"].(float64))
	}
	salt := utils.GetSalt(params["password"].(string))
	user := model.User{
		Name:     params["name"].(string),
		Phone:    params["phone"].(string),
		Salt:     salt,
		Password: utils.Md5(params["password"].(string) + salt),
		RoleId:   int(params["role_id"].(float64)),
		Status:   status,
	}
	result := db.Create(&user)
	return utils.R(result, nil)
}

// UserEdit 修改
func UserEdit(id string, params map[string]interface{}) (bool, interface{}) {
	user := model.User{}
	db.First(&user, id)
	user.Name = params["name"].(string)
	user.Phone = params["phone"].(string)
	user.RoleId = int(params["role_id"].(float64))
	user.Status = int(params["status"].(float64))
	if params["password"] != nil && params["password"] != "" {
		bol, res := SetPwd(params["password"].(string), user.Salt)
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
