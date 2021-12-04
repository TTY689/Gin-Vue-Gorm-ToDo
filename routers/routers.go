package routers

import (
	"github.com/gin-gonic/gin"
	"in_Vue_Gorm_ToDo/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//设置静态文件路径
	r.Static("/static", "static")
	//设置模板文件路径
	r.LoadHTMLGlob("templates/*")
	//设置路由处理
	r.GET("/", controller.IndexHandler)

	v1Group := r.Group("v1")
	{
		v1Group.POST("/todo", controller.CreateToDo)
		v1Group.GET("/todo", controller.GetToDoList)
		v1Group.PUT("/todo:id", controller.UpdateAToDo)
		v1Group.DELETE("/todo:id", controller.DeleteToDo)
	}
	return r

}
