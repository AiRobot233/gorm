package model

import (
	"errors"
	"gin/utils"
	"gorm.io/gorm"
)

//
type Role struct {
	Id        int             `gorm:"column:id" json:"id"`                 //是否可空:NO
	Pid       int             `gorm:"column:pid" json:"pid"`               //是否可空:NO 上级id
	Name      string          `gorm:"column:name" json:"name"`             //是否可空:NO 名称
	Rule      string          `gorm:"column:rule" json:"rule"`             //是否可空:NO 权限
	CreatedAt utils.LocalTime `gorm:"column:created_at" json:"created_at"` //是否可空:NO
	UpdatedAt utils.LocalTime `gorm:"column:updated_at" json:"updated_at"` //是否可空:NO
	DeletedAt *gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`          //是否可空:YES
}

func (*Role) TableName() string {
	return "role"
}

//设置数据体
func RoleSetFromData(role *Role, params map[string]interface{}) {
	role.Pid = int(params["pid"].(float64))
	role.Name = params["name"].(string)
	role.Rule = params["rule"].(string)
}

//删除事件
func (r *Role) BeforeDelete(tx *gorm.DB) (err error) {
	if r.Id == 1 {
		return errors.New("系统账号不允许删除")
	}
	role := Role{}
	result := tx.Model(r).Where("pid = ?", r.Id).First(&role)
	if result.RowsAffected > 0 {
		return errors.New("有子级不允许删除")
	}
	return
}
