package initialize

import (
	"os"
	"server/global"
	"server/task"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Sugar().Infof(msg, keysAndValues...)
}

func (z *ZapLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	z.logger.Sugar().Errorf(msg+": %v", append(keysAndValues, err)...)
}

func NewZapLogger() *ZapLogger {
	return &ZapLogger{logger: global.Log}
}

func InitCron() {
	c := cron.New(cron.WithLogger(NewZapLogger()))

	if err := task.RegisterTasks(c); err != nil {
		global.Log.Error("注册定时任务失败", zap.Error(err))
		os.Exit(0)
	}
}
