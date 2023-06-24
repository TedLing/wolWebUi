package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"wolWebUi/Repository"
	"wolWebUi/Tools"
)

func GetRouter() *gin.Engine {

	r := gin.Default()

	// 浏览界面处理
	r.LoadHTMLFiles("./static/index.html")
	// 注册静态文件:参数1：别名、参数2：当前static文件目录，
	r.Static("static", "./static")
	// 注册路由
	r.GET("/", func(context *gin.Context) {
		context.HTML(http.StatusOK, "index.html", nil)
	})

	//api
	apiGroup := r.Group("/api")

	//dev -api
	devGroup := apiGroup.Group("/dev")

	//查询设备列表
	devGroup.GET("/list", func(c *gin.Context) {
		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, Repository.GetDevList()))
	})

	//新增
	devGroup.POST("/add", func(c *gin.Context) {

		data, err := c.GetRawData()
		if err != nil {
			fmt.Println("获取请求数据失败！")
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		err = Repository.AddDev(data)
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, "[]"))
	})

	//修改
	devGroup.POST("/update", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	//删除
	devGroup.POST("/delete", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	//设备唤醒
	devGroup.POST("/wakeup", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	return r
}
