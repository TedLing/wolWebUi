package main

import (
	"embed"
	"fmt"
	viper "github.com/spf13/viper"
	"strconv"
	"wolWebUi/Tools"
	"wolWebUi/router"
)

var port int

//go:embed  static/*
var content embed.FS

func main() {

	//先连接一下数据库,失败了直接关闭
	fmt.Println("******尝试连接数据库******")
	_ = Tools.GetDB()
	fmt.Println("******连接数据库成功******")

	//引入路由 传入打包的文件
	r := router.GetRouter(content)

	// 启动服务
	strPort := strconv.Itoa(port)
	fmt.Println("尝试开启服务 端口号:" + strPort)
	r.Run("0.0.0.0:" + strPort)
}

// 初始化获取 配置文件
func init() {

	config := viper.New()
	config.AddConfigPath("./")
	config.SetConfigName("wolConf")
	config.SetConfigType("json")

	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("未找到配置文件，端口号默认8899..")
			port = 8899
			return
		} else {
			fmt.Println("获取配置文件出错，端口号默认8899")
			port = 8899
			return
		}
	}

	port = config.GetInt("port")

}
