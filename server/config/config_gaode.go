package config

// Gaode 高德配置
type Gaode struct {
	Enable bool   `json:"enable" yaml:"enable"` // 是否启用
	Key    string `json:"key" yaml:"key"`       // 高德Key
}
