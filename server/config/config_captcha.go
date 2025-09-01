package config

// Captcha 验证码配置
type Captcha struct {
	Height   int     `json:"height" yaml:"height"`       // 高度
	Width    int     `json:"width" yaml:"width"`         // 宽度
	Length   int     `json:"length" yaml:"length"`       // 长度
	MaxSkew  float64 `json:"max_skew" yaml:"max_skew"`   // 最大倾斜度
	DotCount int     `json:"dot_count" yaml:"dot_count"` // 点数
}
