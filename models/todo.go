package models

import "github.com/yolomc/gin_demo/dao"

//Todo Model
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

//InitModel 初始化Model
func InitModel() {
	//模型绑定，自动迁移
	dao.DB.AutoMigrate(&Todo{})
}

//Todo 增删改查

//GetTodoList 获取todo列表数据
func GetTodoList() (todoList []Todo, err error) {
	//var todoList []Todo
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

//GetTodoByID 通过id查询并返回todo对象
func GetTodoByID(id int) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

//AddTodo 新增一条数据
func AddTodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return
}

//UpdateTodo 更新一条数据 by id
func UpdateTodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

//DeleteTodo 删除一条数据 by id
func DeleteTodo(id int) (err error) {
	err = dao.DB.Where("id=?", id).Delete(Todo{}).Error
	return
}
