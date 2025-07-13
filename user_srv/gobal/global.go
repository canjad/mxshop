package gobal

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var (
	// DB 全局数据库连接
	DB *gorm.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer, prefix, flag
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // 彩色打印
		},
	)
	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// schema.NamingStrategy 用于设置表名的命名策略
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名
		},
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("connect to mysql failed, err:%v\n", err)
	}
}
