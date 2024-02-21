package services

import (
	"fmt"
	"gae-cli/config"
	"gae-cli/gsc"
)

// 提供登录api
func Login(username string, password string) bool {
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "Users").Call("login", []string{username, password})
	if !res.GetStatus() {
		println(res.GetMessage())
		return false
	}
	// 获得 token
	result := res.AsObject()
	token := result["token"].(string)
	// 写入 token 到当前 provider
	c := config.NewGaeConfigManager()
	c.SetDefaultProviderToken(token)
	fmt.Printf("%s login ok!\n", username)
	return true
}

// 提供退出api
func Logout() {
	c := config.NewGaeConfigManager()
	c.SetDefaultProviderToken("")
	println("logout ok!")
}
