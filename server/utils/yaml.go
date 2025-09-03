package utils

import (
	"io/fs"
	"os"
	"server/global"

	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

// LoadYAML 读取配置文件内容
func LoadYAML() ([]byte, error) { return os.ReadFile(configFile) }

// SaveYAML 将配置写入配置文件
func SaveYAML() error {
	byteData, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, byteData, fs.ModePerm) // 0777
}
