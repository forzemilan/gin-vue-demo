package common

import (
	"fmt"
	"ginessential/model"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
)

// DB database instance
var DB *gorm.DB

// InitDB 初始化数据库
func InitDB() *gorm.DB {
	driveName := viper.GetString("database.driveName")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")
	database := viper.GetString("database.database")
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	charset := viper.GetString("database.charset")
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
