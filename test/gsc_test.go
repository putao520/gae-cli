package test

import (
	"gae-cli/gsc/services"
	"gae-cli/handle"
	"testing"
	"time"
)

func TestProvider(t *testing.T) {
	p := handle.NewGaeProviderManager()

	// 新增一个provider
	p.AddProvider("test", "http://localhost", "123123")

	// 编辑一个provider
	p.UpdateProviderHost("test", "http://localhost:8080")

	// 删除一个provider
	p.RemoveProvider("test")

	// 新增一个provider
	p.AddProvider("local", "http://127.0.0.1:805", "grapeSoft@")

	// 设置默认provider
	p.DefaultProvider("local")
}

func TestClearProvider(t *testing.T) {
	p := handle.NewGaeProviderManager()
	p.RemoveProvider("local")
}

func TestUserApi(t *testing.T) {
	TestProvider(t)
	services.Login("putao520", "123123")
	// 延迟 5 秒
	time.Sleep(5 * time.Second)
	services.Logout()
	TestClearProvider(t)
}
