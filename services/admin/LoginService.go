package admin

import (
	"gin/model"
	"gin/utils"
)

//登录
func Login(phone string, password string) (bool, string) {
	user := model.User{}
	user.Phone = phone
	result := db.First(&user)
	if result.RowsAffected > 0 {
		p := utils.Md5(password + user.Salt)
		if p == user.Password {
			//登录成功
			token := utils.GetJwt(user)
			return true, token
		} else {
			return false, "密码错误"
		}
	} else {
		return false, "用户未找到"
	}
}
