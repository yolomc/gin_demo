package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yolomc/gin_demo/models"
)

//IndexController 渲染index页面
func IndexController(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

//GetTodoList 获取列表数据
func GetTodoList(c *gin.Context) {
	//定义todo类型切片接收数据
	if todoList, err := models.GetTodoList(); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todoList)
	}
}

//AddTodo 新增一条todo
func AddTodo(c *gin.Context) {

	var todo models.Todo
	c.BindJSON(&todo) //将请求数据与 todo对象绑定

	//插入数据，返回结果
	if err := models.AddTodo(&todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//UpdateTodo 更新一条todo
func UpdateTodo(c *gin.Context) {
	//获取id
	id, ok := getParamsID(c)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	//根据id获取数据
	todo, err := models.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
		return
	}

	//将传入数据绑定到todo对象
	c.BindJSON(todo)

	//更新数据
	if err = models.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

//DeleteTodo 删除一条todo
func DeleteTodo(c *gin.Context) {
	//获取id
	id, ok := getParamsID(c)
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}

	//根据id删除数据
	if err := models.DeleteTodo(id); err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error})
	} else {
		c.JSON(http.StatusOK, gin.H{"ok": "success"})
	}
}

//getParamsID 返回int类型的id参数
func getParamsID(c *gin.Context) (int, bool) {
	idstr, ok := c.Params.Get("id")
	if !ok {
		return 0, false
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		return 0, false
	}

	return id, true
}
