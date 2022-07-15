package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
)

func GetDb() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db
}

//分页
func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		pageSize, _ := strconv.Atoi(pageSize)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
