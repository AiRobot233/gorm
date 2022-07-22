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
	result := db.Table("user").Select([]string{"id", "name", "phone", "created_at"}).Scopes(model.Paginate(page, pageSize)).Scan(&users).Count(&count)
	return utils.R(result, utils.P(users, count))
}

//新增
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

//修改
func UserEdit(id string, params map[string]interface{}) (bool, interface{}) {
	user := model.User{}
	db.First(&user, id)
	user.Name = params["name"].(string)
	user.Phone = params["phone"].(string)
	user.RoleId = int(params["role_id"].(float64))
	user.Status = int(params["status"].(float64))
	if params["password"] != nil && params["password"] != "" {
		err := utils.CheckPasswordLever(params["password"].(string)) //校验密码安全性
		if err != nil {
			return false, err.Error()
		}
		user.Password = utils.Md5(params["password"].(string) + user.Salt)
	}
	result := db.Save(&user)
	return utils.R(result, nil)
}

//删除
func UserDel(id string) (bool, interface{}) {
	user := model.User{}
	result := db.Delete(&user, id)
	return utils.R(result, nil)
}
