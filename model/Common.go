package model

import (
	"gin/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"time"
)

func GetDb() *gorm.DB {
	dsn := utils.GetEnvData("MYSQL_DSN")
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	sqlDB, _ := db.DB()
	sqlDB.SetConnMaxLifetime(time.Second * 60) //设置链接池超时时间
	return db
}

// Paginate 分页
func Paginate(page string, pageSize string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		pageSize, _ := strconv.Atoi(pageSize)
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
