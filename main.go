package main

import (
	"github.com/yolomc/gin_demo/dao"
	"github.com/yolomc/gin_demo/models"
	"github.com/yolomc/gin_demo/routers"
)

func main() {

	err := dao.InitMySQL() //连接数据库
	if err != nil {
		panic(err)
	}
	defer dao.Close() //关闭连接

	models.InitModel() // 初始化Model

	r := routers.SetupRouter() //配置路游

	r.Run() //启动服务
}
