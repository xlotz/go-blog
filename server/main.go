package main

import (
	"go-blog/core"
	"go-blog/global"
	"go-blog/routers"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 初始化日志配置
	global.Log = core.InitLogger()
	// 初始化数据库连接
	global.DB = core.InitGorm()
	// 初始化路由
	router := routers.InitRouter()
	router.Run(global.Config.System.Addr())

}
