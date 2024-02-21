package config

import (
	"encoding/json"
	"gae-cli/file"
	"os"
	"os/user"
)

// 定义配置的授权结构
type GaeConfigProvider struct {
	Host    string `json:"host"`
	Token   string `json:"token"`
	AuthKey string `json:"authKey"`
}

// 定义配置 json 结构
type GaeConfig struct {
	DefaultProvider string                       `json:"defaultProvider"`
	Providers       map[string]GaeConfigProvider `json:"providers"`
}

func buildNewGaeConfig() GaeConfig {
	return GaeConfig{
		DefaultProvider: "gae",
		Providers: map[string]GaeConfigProvider{
			"gae": GaeConfigProvider{
				Host:    "https://gae.putao282.com:805",
				Token:   "",
				AuthKey: "GrapeAuthKey@",
			},
		},
	}
}

func CreateOrUpdateGaeConfig(content GaeConfig) {
	userPathString, err := userprofile()
	if err != nil {
		println("userprofile error")
		return
	}
	configPath := userPathString + "/.gaeconfig"
	createOrUpdateGaeConfigImpl(configPath, content)
}

func createOrUpdateGaeConfigImpl(configPath string, content GaeConfig) {
	file, err := os.Create(configPath)
	if err != nil {
		println("open .gaeconfig file error")
		return
	}
	defer file.Close()
	// content 转换为 json 并写入文件
	encoder := json.NewEncoder(file)
	err = encoder.Encode(content)
	if err != nil {
		println("encode .gaeconfig file error")
		return
	}
}

// 获得 userprofile 目录
func userprofile() (string, error) {
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	var userProfileDir string

	// 如果是Windows系统
	if os.PathSeparator == '\\' {
		// 在Windows中，UserProfile通常位于环境变量USERPROFILE指定的位置
		userProfileDir = os.Getenv("USERPROFILE")
	} else {
		// 在其他系统中，可以使用user包获取当前用户的HomeDir
		userProfileDir = currentUser.HomeDir
	}
	return userProfileDir, nil
}

// 从系统 userprofile 目录下读取配置文件 .gaeconfig
func ReadGaeConfig() *GaeConfig {
	userPathString, err := userprofile()
	if err != nil {
		println("userprofile error")
		return nil
	}
	r := GaeConfig{}
	// 判断配置文件是否存在
	configPath := userPathString + "/.gaeconfig"
	if f, _ := file.PathExists(configPath); f == false {
		// 配置文件不存在创建默认配置文件
		r = buildNewGaeConfig()
		createOrUpdateGaeConfigImpl(configPath, r)
	} else {
		// 读取配置文件
		file, err := os.Open(configPath)
		if err != nil {
			println("open .gaeconfig file error")
			return nil
		}
		defer file.Close()
		// 读取配置文件
		decoder := json.NewDecoder(file)
		err = decoder.Decode(&r)
		if err != nil {
			println("decode .gaeconfig file error")
			return nil
		}
	}
	return &r
}
