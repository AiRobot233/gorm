package admin

import (
	"gin/model"
	"gin/utils"
)

//字典列表
func DictionaryList() (bool, interface{}) {
	var dictionary []*model.DictionaryTree
	result := db.Find(&dictionary)
	return utils.R(result, DictionaryTree(dictionary, 0))
}
