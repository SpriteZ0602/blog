package config

type Qiniu struct {
	Zone          string `json:"zone" yaml:"zone"`                       // 存储区域
	Bucket        string `json:"bucket" yaml:"bucket"`                   // 空间名称
	ImgPath       string `json:"img_path" yaml:"img_path"`               // 图片路径
	AccessKey     string `json:"access_key" yaml:"access_key"`           // AccessKey
	SecretKey     string `json:"secret_key" yaml:"secret_key"`           // SecretKey
	UseHTTPS      bool   `json:"use_https" yaml:"use_https"`             // 是否使用https域名
	UseCdnDomains bool   `json:"use_cdn_domains" yaml:"use_cdn_domains"` // 上传是否使用CDN加速
}
