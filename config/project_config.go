package config

import (
	"encoding/json"
	"os"
)

type GaeProjectConfigInfo struct {
	Provider       string `json:"provider"`
	AppID          string `json:"appId"`
	ServiceID      string `json:"serviceId"`
	DockerRegistry string `json:"dockerRegistry"`
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
	projectConfigPath := path + "/.gae"
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

func createOrUpdateGaeProjectConfigImpl(configPath string, content GaeProjectConfig) {
	// 写入配置文件
	file, err := os.Create(configPath)
	if err != nil {
		println("open .gae file error")
		return
	}
	defer file.Close()
	// content 转换为 json 并写入文件
	encoder := json.NewEncoder(file)
	err = encoder.Encode(content)
	if err != nil {
		println("encode .gae file error")
		return
	}
}

var gaeProjectConfigInstance *GaeProjectConfig

func InitGae() {
	// 读取配置文件
	gaeProjectConfigInstance = ReadGaeProjectConfig()
}

func GetGaeProjectConfig() *GaeProjectConfig {
	if gaeProjectConfigInstance == nil {
		InitGae()
	}
	return gaeProjectConfigInstance
}

func CreateOrUpdateGaeProjectConfig(config GaeProjectConfig) {
	path, err := os.Getwd()
	if err != nil {
		return
	}
	// 读取当前目录下的 .gae 文件(项目配置文件)
	projectConfigPath := path + "/.gae"
	createOrUpdateGaeProjectConfigImpl(projectConfigPath, config)
}
