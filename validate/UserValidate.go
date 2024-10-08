package validate

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
)

// User 数据体最好公开可以外部直接调用
type User struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Phone    string `form:"phone" json:"phone"`
	RoleId   int    `form:"role_id" json:"role_id" binding:"required"`
	Status   int    `form:"status" json:"status"`
	Password string `form:"password" json:"password"`
	UnitId   int    `form:"unit_id" json:"unit_id" binding:"required"`
}

func UserValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	data, _ := io.ReadAll(context.Request.Body)
	// 再重新写回请求体body中
	context.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	var l User
	err := context.ShouldBindWith(&l, binding.JSON)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			message := utils.GetMapValue(errs.Translate(Trans))
			utils.ValidateError(context, message)
			return false
		}
		utils.Error(context, err.Error())
		return false
	}
	// 再重新写回请求体body中
	context.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	return true
}
