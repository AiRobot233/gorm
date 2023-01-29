package validate

import (
	"bytes"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"io/ioutil"
)

// Dictionary 数据体最好公开可以外部直接调用
type Dictionary struct {
	Pid   int    `form:"pid" json:"pid"`
	Name  string `form:"name" json:"name" binding:"required"`
	Value string `form:"value" json:"value" binding:"required"`
	Sort  int    `form:"sort" json:"sort"`
}

func DictionaryValidate(context *gin.Context) bool {
	if err := ValidatorTrans("zh"); err != nil {
		utils.Error(context, err.Error())
		return false
	}
	data, _ := ioutil.ReadAll(context.Request.Body)
	// 再重新写回请求体body中
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	var l Dictionary
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
	context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	return true
}
