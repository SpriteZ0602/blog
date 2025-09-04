package core

import (
	"server/global"
	"server/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

// RunServer 运行服务器
func RunServer() {
	addr := global.Config.System.Addr()
	Router := initialize.InitRouter()

	// 启动服务器
	s := initServer(addr, Router)
	global.Log.Info("服务器启动成功", zap.String("监听地址", addr))
	global.Log.Error(s.ListenAndServe().Error())
}
