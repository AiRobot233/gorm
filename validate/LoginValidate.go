package validate

import (
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type loginValidate struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	var l loginValidate
	err := context.ShouldBindWith(&l, binding.Form)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(context, errs.Translate(Trans))
			return false
		}
	}
	return true
}
