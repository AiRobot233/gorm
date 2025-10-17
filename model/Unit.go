package model

import (
	"errors"
	"gin/utils"
	"gin/validate"
	"gorm.io/gorm"
)

type Unit struct {
	Id         int              `gorm:"column:id" json:"id,omitempty"`     //是否可空:NO
	Pid        int              `gorm:"column:pid" json:"pid"`             //是否可空:NO 上级id
	Sort       int              `gorm:"column:sort" json:"sort"`           //是否可空:NO
	Name       string           `gorm:"column:name" json:"name,omitempty"` //是否可空:NO 名称
	Type       string           `gorm:"column:type" json:"type,omitempty"` //是否可空:NO 类型
	IsUnit     int              `gorm:"is_unit" json:"is_unit" `
	IsRegister int              `gorm:"is_register" json:"is_register" `
	CheckOrg   int              `gorm:"check_org" json:"check_org"`
	CreatedAt  *utils.LocalTime `gorm:"column:created_at" json:"created_at,omitempty"` //是否可空:NO
	UpdatedAt  *utils.LocalTime `gorm:"column:updated_at" json:"updated_at,omitempty"` //是否可空:NO
	DeletedAt  gorm.DeletedAt   `gorm:"column:deleted_at" json:"-"`                    //是否可空:NO
}

func (*Unit) TableName() string {
	return "unit"
}

// UnitTree 定义树状结构体
type UnitTree struct {
	Unit
	Child []*UnitTree `gorm:"-" json:"children"`
}

// UnitSetFromData 设置数据
func (r *Unit) UnitSetFromData(params validate.Unit) {
	r.Pid = *params.Pid
	r.Name = params.Name
	r.IsUnit = params.IsUnit
	r.IsRegister = params.IsRegister
	r.Sort = params.Sort
	r.Type = params.Type
	r.CheckOrg = params.CheckOrg
}

func (r *Unit) BeforeCreate(tx *gorm.DB) (err error) {
	// 单位创建时自动创建后台账号
	if r.IsUnit == 1 {
		password := "Aa@112233"
		salt := utils.GetSalt(password)
		tx.Create(&User{
			Name:     r.Name,
			Salt:     salt,
			Password: utils.Md5(password + salt),
			RoleId:   6,
			Status:   1,
		})
	}
	return
}

// BeforeSave 新增修改事件
func (r *Unit) BeforeSave(tx *gorm.DB) (err error) {
	if r.Id > 0 && r.Id == r.Pid {
		return errors.New("上级单位不能选择自己")
	}
	unit := Unit{}
	result := tx.Model(r).Where("id != ? AND name = ?", r.Id, r.Name).First(&unit)
	if result.RowsAffected > 0 {
		return errors.New("单位已存在")
	}
	return
}
