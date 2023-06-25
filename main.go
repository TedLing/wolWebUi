package main

import (
	"wolWebUi/Tools"
	"wolWebUi/router"
)

func main() {

	//先连接一下数据库,失败了直接关闭
	_ = Tools.GetDB()

	//引入路由
	r := router.GetRouter()

	// 启动服务
	r.Run()
}
