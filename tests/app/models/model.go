package models

import (
	"goblog/pkg/logger"
	"goblog/pkg/types"

	gormlogger "gorm.io/gorm/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() *gorm.DB {
	var err error

	config := mysql.New(mysql.Config{
		DSN: "wcs:yanru520@tcp(127.0.0.1:3306)/goblog?charset=utf8&parseTime=True&loc=Local",
	})

	// 准备数据库连接池
	DB, err = gorm.Open(config, &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Info),
	})
	logger.LogError(err)

	return DB
}

type BaseModel struct {
	ID uint64
}

func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
