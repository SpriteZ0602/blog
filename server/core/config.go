package core

import (
	"log"
	"server/config"
	"server/utils"

	"gopkg.in/yaml.v3"
)

// InitConfig 初始化配置
func InitConfig() *config.Config {
	c := &config.Config{}
	yamlConfig, err := utils.LoadYAML()
	if err != nil {
		log.Fatalf("读取配置文件失败: %v", err)
	}
	if err = yaml.Unmarshal(yamlConfig, c); err != nil {
		log.Fatalf("解析配置文件失败: %v", err)
	}
	return c
}
