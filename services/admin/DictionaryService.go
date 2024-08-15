package admin

import (
	"gin/model"
	"gin/utils"
	"gin/validate"
)

// DictionaryList 字典列表
func DictionaryList() (bool, any) {
	var dictionary []*model.DictionaryTree
	result := db.Find(&dictionary)
	return utils.R(result, DictionaryTree(dictionary, 0))
}

// DictionaryAdd 字典添加
func DictionaryAdd(params validate.Dictionary) (bool, any) {
	dictionary := model.Dictionary{}
	dictionary.DictionarySetFromData(params)
	result := db.Create(&dictionary)
	return utils.R(result, nil)
}

// DictionaryEdit 字典修改
func DictionaryEdit(id string, params validate.Dictionary) (bool, any) {
	dictionary := model.Dictionary{}
	res := db.First(&dictionary, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	dictionary.DictionarySetFromData(params)
	result := db.Save(&dictionary)
	return utils.R(result, nil)
}

// DictionaryDel 字典删除
func DictionaryDel(id string) (bool, any) {
	dictionary := model.Dictionary{}
	res := db.First(&dictionary, id)
	if res.Error != nil {
		return false, res.Error.Error()
	}
	result := db.Delete(&dictionary)
	return utils.R(result, nil)
}

// UnitDictionary 获取字典数据（不鉴权）
func UnitDictionary(name string) (bool, any) {
	dictionary := model.Dictionary{}
	result := db.Where("name = ?", name).First(&dictionary)
	if result.Error != nil {
		return false, result.Error.Error()
	}
	var tree []*model.DictionaryTree
	res := db.Where("pid = ?", dictionary.Id).Find(&tree)
	return utils.R(res, DictionaryTree(tree, dictionary.Id))
}

// DictionarySelect 字典下拉
func DictionarySelect(value string) (bool, any) {
	var dictionary []*model.DictionaryTree
	result := db.Find(&dictionary)
	pid := 0
	if value != "" {
		var a model.Dictionary
		_ = db.Where("value", value).Find(&a)
		pid = a.Id
	}
	return utils.R(result, DictionaryTree(dictionary, pid))
}
