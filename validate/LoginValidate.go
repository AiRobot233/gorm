package validate

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

//数据体最好公开可以外部直接调用
type Login struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	var l Login
	err := context.ShouldBindWith(&l, binding.Form)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(context, errs.Translate(Trans))
			return false
		}
	}
	return true
}
