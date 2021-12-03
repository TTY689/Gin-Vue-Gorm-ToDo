package routers

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//设置静态文件路径
	r.Static("/static", "static")
	//设置模板文件路径
	r.LoadHTMLGlob("templates/*")
	//设置路由处理
	r.GET("/")

}
