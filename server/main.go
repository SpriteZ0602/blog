package main

import (
	"server/core"
	"server/global"
)

func main() {
	// 1. 初始化配置
	global.Config = core.InitConfig()
	// 2. 初始化日志
	global.Log = core.InitLogger()

	// 3. 运行服务器
	core.RunServer()
}
