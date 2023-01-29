package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// DictionaryList 字典列表
func DictionaryList() (bool, interface{}) {
	var dictionary []*model.DictionaryTree
	result := db.Find(&dictionary)
	return utils.R(result, DictionaryTree(dictionary, 0))
}

// DictionaryAdd 字典添加
func DictionaryAdd(params validate.Dictionary) (bool, interface{}) {
	dictionary := model.Dictionary{}
	dictionary.DictionarySetFromData(params)
	result := db.Create(&dictionary)
	return utils.R(result, nil)
}

// DictionaryEdit 字典修改
func DictionaryEdit(id string, params validate.Dictionary) (bool, interface{}) {
	dictionary := model.Dictionary{}
	db.First(&dictionary, id)
	dictionary.DictionarySetFromData(params)
	result := db.Save(&dictionary)
	return utils.R(result, nil)
}

// DictionaryDel 字典删除
func DictionaryDel(id string) (bool, interface{}) {
	dictionary := model.Dictionary{}
	res := db.Where("id = ?", id).First(&dictionary)
	if res.RowsAffected == 0 {
		return false, "数据不存在"
	}
	result := db.Delete(&dictionary)
	return utils.R(result, nil)
}

// UnitDictionary 获取字典数据（不鉴权）
func UnitDictionary(name string) (bool, interface{}) {
	dictionary := model.Dictionary{}
	result := db.Where("name = ?", name).First(&dictionary)
	if result.RowsAffected == 0 {
		return false, "数据不存在"
	}
	var tree []*model.DictionaryTree
	res := db.Where("pid = ?", dictionary.Id).Find(&tree)
	return utils.R(res, DictionaryTree(tree, dictionary.Id))
}

// DictionarySelect 字典下拉
func DictionarySelect() (bool, interface{}) {
	var dictionary []*model.DictionaryTree
	result := db.Find(&dictionary)
	return utils.R(result, DictionaryTree(dictionary, 0))
}
