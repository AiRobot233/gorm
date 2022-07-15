package admin

import (
	"gin/model"
	"gin/utils"
)

//定义列表展示字段
type listUser struct {
	Id        int             `json:"id"`
	Name      string          `json:"name"`
	Phone     string          `json:"phone"`
	CreatedAt utils.LocalTime `json:"created_at"`
}

//列表
func UserList(page string, pageSize string) (bool, interface{}) {
	var users []listUser //定义表结构
	var count int64
	result := db.Table("user").Select([]string{"id", "name", "created_at"}).Scopes(model.Paginate(page, pageSize)).Scan(&users).Count(&count)
	if result.Error != nil {
		return false, result.Error.Error()
	} else {
		return true, utils.P(users, count)
	}
}

//新增
func UserAdd(password string, name string, phone string) (bool, interface{}) {
	salt := utils.GetSalt(password)
	user := model.User{
		Name:     name,
		Phone:    phone,
		Salt:     salt,
		Password: utils.Md5(password + salt),
	}
	result := db.Create(&user)
	if result.Error != nil {
		return false, result.Error.Error()
	} else {
		return true, nil
	}
}

//修改
func UserEdit(id string, params map[string]string) (bool, interface{}) {
	user := model.User{}
	db.First(&user, id)
	user.Name = params["name"]
	user.Phone = params["phone"]
	if params["password"] != "" {
		user.Password = utils.Md5(params["password"] + user.Salt)
	}
	result := db.Save(&user)
	if result.Error != nil {
		return false, result.Error.Error()
	} else {
		return true, nil
	}
}

//删除
func UserDel(id string) (bool, interface{}) {
	user := model.User{}
	result := db.Delete(&user, id)
	if result.Error != nil {
		return false, result.Error.Error()
	} else {
		return true, nil
	}
}
