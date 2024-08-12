package model

import (
	"errors"
	"gin/validate"
	"gorm.io/gorm"
)

type Dictionary struct {
	Id    int    `gorm:"column:id" json:"id"`       //是否可空:NO
	Pid   int    `gorm:"column:pid" json:"pid"`     //是否可空:NO
	Name  string `gorm:"column:name" json:"name"`   //是否可空:NO 名称
	Value string `gorm:"column:value" json:"value"` //是否可空:NO 值
	Sort  int    `gorm:"column:sort" json:"sort"`   //是否可空:NO 排序
}

func (*Dictionary) TableName() string {
	return "dictionary"
}

// DictionaryTree 定义树状结构体
type DictionaryTree struct {
	Dictionary
	Child []*DictionaryTree `gorm:"-" json:"children"`
}

// DictionarySetFromData 设置数据
func (r *Dictionary) DictionarySetFromData(params validate.Dictionary) {
	r.Pid = params.Pid
	r.Name = params.Name
	r.Value = params.Value
	r.Sort = params.Sort
}

// BeforeDelete 删除事件
func (r *Dictionary) BeforeDelete(tx *gorm.DB) (err error) {
	dictionary := Dictionary{}
	result := tx.Model(r).Where("pid = ?", r.Id).First(&dictionary)
	if result.RowsAffected > 0 {
		return errors.New("有子级不允许删除")
	}
	return
}

// BeforeSave 新增修改事件
func (r *Dictionary) BeforeSave(tx *gorm.DB) (err error) {
	dictionary := Dictionary{}
	result := tx.Model(r).Where("id != ? AND name = ? AND pid = ?", r.Id, r.Name, r.Pid).First(&dictionary)
	if result.RowsAffected > 0 {
		return errors.New("名称已存在")
	}
	return
}
