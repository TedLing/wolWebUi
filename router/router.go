package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
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
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		err = Repository.UpdateDev(data)
		if err != nil {
			fmt.Println(Tools.GetFailMsg(err.Error()))
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, "[]"))
	})

	//删除
	devGroup.POST("/delete", func(c *gin.Context) {
		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		err = Repository.DelDev(data)
		if err != nil {
			fmt.Println(Tools.GetFailMsg(err.Error()))
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, "[]"))
	})

	//设备唤醒
	devGroup.POST("/wakeup", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	//接口列表
	apiGroup.GET("/interfaces", func(c *gin.Context) {

		//获取所有的接口
		ifaces, err := net.Interfaces()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		interfaces := make(map[int]string)
		for _, iface := range ifaces {
			interfaces[len(interfaces)+1] = iface.Name
		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, interfaces))
	})

	return r
}
