package config

// ES 配置
type ES struct {
	URL            string `json:"url" yaml:"url"`                           // Elasticsearch URL
	Username       string `json:"username" yaml:"username"`                 // Elasticsearch Username
	Password       string `json:"password" yaml:"password"`                 // Elasticsearch Password
	IsConsolePrint bool   `json:"is_console_print" yaml:"is_console_print"` // 是否在控制台打印日志
}
