package validate

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
)

// Login 数据体最好公开可以外部直接调用
type Login struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func LoginValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	data, _ := io.ReadAll(context.Request.Body)
	// 再重新写回请求体body中
	context.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	var l Login
	err := context.ShouldBindWith(&l, binding.JSON)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			utils.ValidateError(context, errs.Translate(Trans))
			return false
		}
		utils.Error(context, err.Error())
		return false
	}
	// 再重新写回请求体body中
	context.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	return true
}
