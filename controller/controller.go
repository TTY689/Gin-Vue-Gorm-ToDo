package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"in_Vue_Gorm_ToDo/models"
	"net/http"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateToDo(c *gin.Context) {
	var todo models.ToDo
	err := c.Bind(&todo)
	if err != nil {
		fmt.Println("表单信息读取错误")
		return
	}

	err = models.CreateAToDo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func GetToDoList(c *gin.Context) {

}
