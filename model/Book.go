package model

//
type Book struct {
	Id     int     `gorm:"column:id" json:"id"`           //是否可空:NO
	Name   *string `gorm:"column:name" json:"name"`       //是否可空:YES
	UserId *int    `gorm:"column:user_id" json:"user_id"` //是否可空:YES
}

func (*Book) TableName() string {
	return "book"
}
