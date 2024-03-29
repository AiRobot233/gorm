package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// Login 登录
func Login(params validate.Login) (bool, any) {
	name := params.Name
	password := params.Password
	user := model.User{}
	result := db.Where("(`phone` = ? OR `name` = ?)", name, name).First(&user)
	if result.RowsAffected > 0 {
		//判断用户是否禁用
		if user.Status == 1 {
			p := utils.Md5(password + user.Salt)
			if p == user.Password {
				//登录成功
				bol, token := utils.GetJwt(user)
				return bol, token
			} else {
				return false, "密码错误"
			}
		} else {
			return false, "用户已被禁用"
		}
	} else {
		return false, "用户未找到"
	}
}
