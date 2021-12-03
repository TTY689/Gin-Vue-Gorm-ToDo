package main

import (
	"fmt"
	"in_Vue_Gorm_ToDo/dao"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		fmt.Print("数据库初始化错误")
		panic(err)
	}

	fmt.Println("数据库链接成功")
}
