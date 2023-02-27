package service

import (
	"bubble/dao"
	"bubble/model"
)

// 业务逻辑
// TODO 增删改查
func CretaeOneTodo(todo *model.Todo) (err error) {

	err = dao.DB.Create(&todo).Error
	return
}

func GetTodoAll() (todoList []*model.Todo, err error) {

	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

func GetTodoOne(id string) (todo *model.Todo, err error) {

	todo = new(model.Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func ModifyTodo(todo *model.Todo) (err error) {

	err = dao.DB.Save(todo).Error
	return
}

func DeleteTodo(id string) (err error) {

	err = dao.DB.Where("id=?", id).Delete(model.Todo{}).Error
	return
}
