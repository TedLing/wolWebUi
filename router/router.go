package router

import (
	"embed"
	"encoding/json"
	"fmt"
	wakeonlan "github.com/ahmetozer/wakeonlan/share"
	"github.com/gin-gonic/gin"
	"html/template"
	"io/fs"
	"net"
	"net/http"
	"strings"
	"wolWebUi/Repository"
	"wolWebUi/Tools"
)

func GetRouter(content embed.FS) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// 浏览界面处理
	r.SetHTMLTemplate(template.Must(template.New("").ParseFS(content, "static/*.html")))
	//r.LoadHTMLFiles("./static/index.html")
	// 注册静态文件:参数1：别名、参数2：当前static文件目录，
	//r.Static("static", "./static")
	fp, _ := fs.Sub(content, "static")
	r.StaticFS("static", http.FS(fp))

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

		data, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		dev := make(map[string]string)
		err = json.Unmarshal(data, &dev)
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		mac, err := net.ParseMAC(dev["mac"])
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}
		packet := wakeonlan.MagicPacket{HWAddr: mac, Device: dev["interface"], IPAddr: dev["ipaddress"], Port: dev["port"]}
		err = packet.SendMagicPacket()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, "[]"))
	})

	//接口列表
	apiGroup.GET("/interfaces", func(c *gin.Context) {

		//获取所有的接口
		ifaces, err := net.Interfaces()
		if err != nil {
			c.JSON(http.StatusOK, Tools.GetFailMsg(err.Error()))
			return
		}

		interfaces := make(map[string]string)
		for _, iface := range ifaces {
			var ips string
			addrs, err := iface.Addrs()
			if err != nil {
				fmt.Println("get addrs err:", err)
				continue
			}
			for _, addr := range addrs {
				//只需要 ipv4 地址
				if strings.Contains(addr.String(), ".") {
					ips += addr.String()
				}
			}

			if len(ips) > 0 {
				interfaces[iface.Name] = "[" + ips + "]"
			}

		}

		c.JSON(http.StatusOK, Tools.GetSuccMsg(999, interfaces))
	})

	return r
}
