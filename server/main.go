package main

import (
	"server/core"
	"server/flags"
	"server/global"
	"server/initialize"
)

func main() {
	// 1. 初始化配置
	global.Config = core.InitConfig()
	// 2. 初始化日志
	global.Log = core.InitLogger()
	// 3. 初始化Redis
	global.Redis = initialize.ConnectRedis()
	// 4. 初始化数据库
	global.DB = initialize.InitGorm()
	// 5. 连接ES
	global.ESClient = initialize.ConnectES()
	// 6. 其他初始化
	initialize.OtherInit()

	// 7. 程序结束前关闭数据库连接
	defer global.Redis.Close()

	// 8. 解析命令行参数
	flags.InitFlag()

	// 9. 初始化定时任务
	initialize.InitCron()

	// 10. 启动服务器
	core.RunServer()
}
