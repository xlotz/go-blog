package main

import (
	"fmt"
	"go-blog/core"
	"go-blog/global"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	fmt.Println(global.Config)

	// 初始化日志配置
	global.Log = core.InitLogger()
	global.Log.Warn("aaaa")
	global.Log.Info("bbbb")
	global.Log.Error("cccc")

	// 初始化数据库连接
	global.DB = core.InitGorm()
	fmt.Println(global.DB)

}
