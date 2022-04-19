package main

import (
	"github.com/zituocn/gow"
	"github.com/zituocn/samblog/conn"
	"github.com/zituocn/samblog/router"
)

func init() {
	// init 日志
	conn.InitLog()

	// init mysql
	conn.InitMySQL()
}

func main() {
	r := gow.Default()

	//设置配置文件信息
	r.SetAppConfig(gow.GetAppConfig())

	//设置静态资源目录
	r.Static("/static", "static")

	//路由设置
	router.PageRouter(r)

	//运行http服务
	r.Run()
}
