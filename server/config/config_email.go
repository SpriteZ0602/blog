package config

// Email 邮件配置
type Email struct {
	Host     string `json:"host" yaml:"host"`         // SMTP服务器地址
	Port     int    `json:"port" yaml:"port"`         // SMTP服务器端口
	From     string `json:"from" yaml:"from"`         // 发件人邮箱
	Nickname string `json:"nickname" yaml:"nickname"` // 发件人昵称
	Secret   string `json:"secret" yaml:"secret"`     // 发件人邮箱授权码
	IsSSL    bool   `json:"is_ssl" yaml:"is_ssl"`     // 是否启用SSL
}
