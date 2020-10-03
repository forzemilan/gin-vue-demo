package common

import (
	"fmt"
	"ginessential/model"

	"github.com/jinzhu/gorm"
)

// DB database instance
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	driveName := "mysql"
	host := "localhost"
	port := "3306"
	database := "ginvuedemo"
	username := "user"
	password := "user"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(driveName, args)
	if err != nil {
		panic("failed to connect databse, err:" + err.Error())
	}
	//! 创建数据表
	db.AutoMigrate(&model.User{})
	//! 关联数据库
	DB = db
	return db
}

// GetDB xx
func GetDB() *gorm.DB {
	return DB
}
