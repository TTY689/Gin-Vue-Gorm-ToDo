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
	todolist, err := models.GetAllToDo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateAToDo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "无效ID",
		})
		return
	}
	todo, err := models.GetAToDo(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	}
	_ = c.BindJSON(&todo)
	err = models.UpdateAToDo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteToDo(c *gin.Context) {
	id, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{
			"error": "id非法",
		})
	}
	if err := models.DeleteAToDo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{})
	}
}
