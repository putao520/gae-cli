package config

import (
	"encoding/json"
	"os"
)

type GaeProjectConfigInfo struct {
	Provider  string `json:"provider"`
	AppID     string `json:"appId"`
	ServiceID string `json:"serviceId"`
}

type GaeProjectConfig struct {
	Project GaeProjectConfigInfo `json:"project"`
}

func ReadGaeProjectConfig() *GaeProjectConfig {
	path, err := os.Getwd()
	if err != nil {
		return nil
	}
	// 读取当前目录下的 .gae 文件(项目配置文件)
	projectConfigPath := path + "/.gaeshell"
	if _, err := os.Stat(projectConfigPath); os.IsExist(err) {
		println(".gae file not found")
		return nil
	}
	var projectConfig, fileErr = os.OpenFile(projectConfigPath, os.O_RDONLY, 0666)
	if fileErr != nil {
		println(".gae file can't open")
		return nil
	}
	defer projectConfig.Close()
	// 读取文件内容
	decoder := json.NewDecoder(projectConfig)
	var config GaeProjectConfig
	err = decoder.Decode(&config)
	if err != nil {
		println("decode .gae file error")
		return nil
	}
	return &config
}

var gaeProjectConfigInstance *GaeProjectConfig

func InitGae() {
	// 读取配置文件
	cfg := ReadGaeProjectConfig()
	if cfg == nil {
		println("ReadGaeProjectConfig error")
	}
	gaeProjectConfigInstance = cfg
}

func GetGaeProjectConfig() *GaeProjectConfig {
	if gaeProjectConfigInstance == nil {
		InitGae()
	}
	return gaeProjectConfigInstance
}
