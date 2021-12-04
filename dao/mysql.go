package dao

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	//稍后版本将更新ini配置文件
	//"github.com/go-ini/ini"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {
	//ini配置待定
	dsn := "czy:密码@tcp(8.136.224.253)/Go?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn}), &gorm.Config{})
	return DB.Error
}

//最新的Gorm已经删除了Close方法
//func Close(){
//	db.
//}
