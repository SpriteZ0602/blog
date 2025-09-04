package initialize

import (
	"server/global"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitGorm() *gorm.DB {
	mysqlConfig := global.Config.Mysql

	db, err := gorm.Open(mysql.Open(mysqlConfig.GetDsn()), &gorm.Config{
		Logger: logger.Default.LogMode(mysqlConfig.GetLogLevel()),
	})

	if err != nil {
		global.Log.Error("连接MySQL异常", zap.Error(err))
	}

	sqlDB, err := db.DB()
	if err != nil {
		global.Log.Error("获取数据库对象异常", zap.Error(err))
	}
	sqlDB.SetMaxIdleConns(mysqlConfig.MaxIdleConns)
	sqlDB.SetMaxOpenConns(mysqlConfig.MaxOpenConns)

	return db
}
