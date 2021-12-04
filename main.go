package main

import (
	"fmt"
	"in_Vue_Gorm_ToDo/dao"
	"in_Vue_Gorm_ToDo/models"
	"in_Vue_Gorm_ToDo/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		fmt.Print("数据库初始化错误")
		panic(err)
	}
	fmt.Println("数据库链接成功")
	err = dao.DB.AutoMigrate(&models.ToDo{})
	if err != nil {
		fmt.Print("数据库迁移失败")
		return
	}
	r := routers.SetupRouter()
	err = r.Run()
	if err != nil {
		fmt.Print("端口启动&运行服务失败")
		return
	}
}
