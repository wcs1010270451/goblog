package model

import (
	"goblog/pkg/logger"

	"gorm.io/gorm"
	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "root:secret@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{})
	logger.LogError(err)

	return DB

}
