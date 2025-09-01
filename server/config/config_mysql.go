package config

import (
	"strconv"
	"strings"

	"gorm.io/gorm/logger"
)

// Mysql 配置
type Mysql struct {
	Host         string `json:"host" yaml:"host"`                     // 服务器地址
	Port         int    `json:"port" yaml:"port"`                     // 端口
	Config       string `json:"config" yaml:"config"`                 // 高级配置
	DBName       string `json:"dbname" yaml:"dbname"`                 // 数据库名
	Username     string `json:"username" yaml:"username"`             // 数据库用户名
	Password     string `json:"password" yaml:"password"`             // 数据库密码
	MaxIdleConns int    `json:"max_idle_conns" yaml:"max_idle_conns"` // 空闲连接数
	MaxOpenConns int    `json:"max_open_conns" yaml:"max_open_conns"` // 最大连接数
	LogMode      string `json:"log_mode" yaml:"log_mode"`             // 日志模式
}

// GetDsn 获取DSN
func (m Mysql) GetDsn() string {
	return m.Username + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.DBName + "?" + m.Config
}

// GetLogLevel 获取日志级别
func (m Mysql) GetLogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	default:
		return logger.Info
	}
}
