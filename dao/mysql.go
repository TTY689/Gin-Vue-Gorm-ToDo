package dao

import (
	"fmt"
	"gopkg.in/ini.v1"
	//"github.com/go-ini/ini"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() (err error) {

	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("配置文件出错")
		return err
	}

	username := cfg.Section("mysql").Key("username").String()
	password := cfg.Section("mysql").Key("password").String()
	ip := cfg.Section("mysql").Key("ip").String()
	database := cfg.Section("mysql").Key("database").String()
	var dsn string
	dsn = username + ":" + password + "@tcp(" + ip + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := "czy:528429@tcp(8.136.224.253)/Go?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn}), &gorm.Config{})
	return DB.Error
}
