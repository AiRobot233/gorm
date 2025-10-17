package validate

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io"
)

// Unit 数据体最好公开可以外部直接调用
type Unit struct {
	Pid        *int   `form:"pid" json:"pid"  binding:"required"`
	Name       string `form:"name" json:"name" binding:"required"`
	IsUnit     int    `form:"is_unit" json:"is_unit" binding:"required"`
	IsRegister int    `form:"is_register" json:"is_register"`
	Sort       int    `form:"sort" json:"sort"`
	Type       string `form:"type" json:"type"`
	CheckOrg   int    `form:"check_org" json:"check_org"`
}

func UnitValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	data, _ := io.ReadAll(context.Request.Body)
	// 再重新写回请求体body中
	context.Request.Body = io.NopCloser(bytes.NewBuffer(data))
	var l Unit
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
