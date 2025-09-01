package config

// QQ 配置
type QQ struct {
	Enabled     bool   `json:"enabled" yaml:"enabled"`
	AppID       string `json:"app_id" yaml:"app_id"`
	AppKey      string `json:"app_key" yaml:"app_key"`
	RedirectURI string `json:"redirect_uri" yaml:"redirect_uri"`
}

func (qq QQ) QQLoginURL() string {
	return "https://graph.qq.com/oauth2.0/authorize?response_type=code&client_id=" + qq.AppID + "&redirect_uri=" + qq.RedirectURI + "&state=STATE"
}
