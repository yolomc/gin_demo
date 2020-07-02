package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/yolomc/gin_demo/controller"
)

//SetupRouter 配置路游并返回gin框架对象
func SetupRouter() *gin.Engine {
	//加载gin框架
	r := gin.Default()
	//指定静态文件路径
	r.Static("/static", "static")
	//指定html文件路径
	r.LoadHTMLGlob("templates/*")
	//获取get请求，渲染并返回index页面内容
	r.GET("/", controller.IndexController)

	//v1 路游下的增删改查操作
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/todo", controller.GetTodoList)
		v1Group.POST("/todo", controller.AddTodo)
		v1Group.PUT("/todo/:id", controller.UpdateTodo)
		v1Group.DELETE("/todo/:id", controller.DeleteTodo)
	}
	return r
}
