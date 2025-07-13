package main

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/hex"

	"github.com/anaskhan96/go-password-encoder"
	"io"
)

// 生成md5加密
func genMd5(code string) string {
	Md5 := md5.New()
	_, _ = io.WriteString(Md5, code)
	return hex.EncodeToString(Md5.Sum(nil))
}
func main() {

	//dsn := "root:123456@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags), // io.Writer, prefix, flag
	//	logger.Config{
	//		SlowThreshold: time.Second, // 慢 SQL 阈值
	//		LogLevel:      logger.Info, // Log level
	//		Colorful:      true,        // 彩色打印
	//	},
	//)
	////全局模式
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	//	// schema.NamingStrategy 用于设置表名的命名策略
	//	NamingStrategy: schema.NamingStrategy{
	//		SingularTable: true, // 使用单数表名
	//	},
	//	Logger: newLogger,
	//})
	//if err != nil {
	//	log.Fatalf("connect to mysql failed, err:%v\n", err)
	//}
	////表迁移 自动创建表结构
	//_ = db.AutoMigrate(&model.User{})
	//	//fmt.Println(genMd5("xxx_123456")) // 输出: e10adc3949ba59abbe56e057f20f883e
	//	//salt, encodedPwd := password.Encode("generic password", nil)
	//	//check := password.Verify("generic password", salt, encodedPwd, nil)
	//	//fmt.Println(check) // true
	//
	//	// Using custom options
	options := &password.Options{16, 10000, 32, sha512.New}

	salt, encodedPwd := password.Encode("generic password", options)
	newPassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	fmt.Println(newPassword)
	//	fmt.Println(len(newPassword))
	//	//解析出 salt 和 encodedPwd
	//	passwordInfo := strings.Split(newPassword, "$")
	//	fmt.Println(passwordInfo)
	//	check := password.Verify("generic password", passwordInfo[2], passwordInfo[3], options)
	//	fmt.Println(check) // true
}
