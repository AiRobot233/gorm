package model

//
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

//定义树状结构体
type DictionaryTree struct {
	Dictionary
	Child []*DictionaryTree `gorm:"-" json:"children"`
}
