package models

import "in_Vue_Gorm_ToDo/dao"

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateAToDo(todo *ToDo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllToDo() (todoList *[]ToDo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetAToDo(id string) (todo *ToDo, err error) {
	todo = new(ToDo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateAToDo(todo *ToDo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteAToDo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(ToDo{}).Error
	return
}
