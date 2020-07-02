package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DB 全局变量
	DB *gorm.DB
)

//InitMySQL 连接数据库
func InitMySQL() (err error) {
	dsn := "root:root1234@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	//连接数据库
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

//Close 关闭数据库连接
func Close() {
	DB.Close()
}
