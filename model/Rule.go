package model

import (
	"errors"
	"gorm.io/gorm"
)

//
type Rule struct {
	Id     int     `gorm:"column:id" json:"id,omitempty"`         //是否可空:NO
	Pid    int     `gorm:"column:pid" json:"pid"`                 //是否可空:NO 上级id
	Name   string  `gorm:"column:name" json:"name,omitempty"`     //是否可空:NO 规则名称
	Type   string  `gorm:"column:type" json:"type,omitempty"`     //是否可空:NO 类型
	Method *string `gorm:"column:method" json:"method"`           //是否可空:YES 请求类型
	Router string  `gorm:"column:router" json:"router,omitempty"` //是否可空:NO 地址/路由
	Sort   int     `gorm:"column:sort" json:"sort"`               //是否可空:NO 排序
}

//定义树状结构体
type RuleTree struct {
	Rule
	Child []*RuleTree `gorm:"-" json:"children"`
}

func (*Rule) TableName() string {
	return "rule"
}

//设置数据
func (r *Rule) RuleSetFromData(params map[string]interface{}) {
	r.Pid = int(params["pid"].(float64))
	r.Name = params["name"].(string)
	r.Type = params["type"].(string)
	r.Router = params["router"].(string)
	if params["sort"] != nil {
		r.Sort = int(params["sort"].(float64))
	}
	if params["method"] != nil {
		method := params["method"].(string)
		r.Method = &method
	}
}

//删除事件
func (r *Rule) BeforeDelete(tx *gorm.DB) (err error) {
	rule := Rule{}
	result := tx.Model(r).Where("pid = ?", r.Id).First(&rule)
	if result.RowsAffected > 0 {
		return errors.New("有子级不允许删除")
	}
	return
}
