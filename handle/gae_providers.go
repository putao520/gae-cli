package handle

import (
	"gae-cli/config"
	"strings"
)

/**
提供 增删改查 应用引擎提供的服务
*/

type GaeProviderManager struct {
	Config *config.GaeConfig
}

func NewGaeProviderManager() *GaeProviderManager {
	cfg := config.ReadGaeConfig()
	if cfg == nil {
		println("ReadGaeConfig error")
	}
	return &GaeProviderManager{
		Config: cfg,
	}
}

// 新增一个服务提供者
func (g *GaeProviderManager) AddProvider(name string, host string, authKey string) {
	g.Config.Providers[name] = config.GaeConfigProvider{
		Host:    host,
		Token:   "",
		AuthKey: authKey,
	}
	config.CreateOrUpdateGaeConfig(*g.Config)
}

// 删除一个服务提供者
func (g *GaeProviderManager) RemoveProvider(name string) {
	// 如果 name 是默认服务提供者, 则删除后需要重新设置默认服务提供者
	if g.Config.DefaultProvider == name {
		g.Config.DefaultProvider = "gae"
	}
	delete(g.Config.Providers, name)
	config.CreateOrUpdateGaeConfig(*g.Config)
}

// 设置默认服务提供者
func (g *GaeProviderManager) DefaultProvider(name string) {
	current := g.Config.Providers[name]
	if current.Host == "" {
		println("provider not found")
		return
	}
	g.Config.DefaultProvider = name
	config.CreateOrUpdateGaeConfig(*g.Config)
}

// 更新一个服务提供者
func (g *GaeProviderManager) UpdateProviderHost(name string, host string) {
	current := g.Config.Providers[name]
	if current.Host == "" {
		println("provider not found")
		return
	}
	g.Config.Providers[name] = config.GaeConfigProvider{
		Host:    host,
		Token:   current.Token,
		AuthKey: current.AuthKey,
	}
	config.CreateOrUpdateGaeConfig(*g.Config)
}

// 更新一个服务提供者
func (g *GaeProviderManager) UpdateProviderAuthKey(name string, authKey string) {
	current := g.Config.Providers[name]
	if current.Token == "" {
		println("provider not found")
		return
	}
	g.Config.Providers[name] = config.GaeConfigProvider{
		Host:    current.Host,
		Token:   current.Token,
		AuthKey: authKey,
	}
	config.CreateOrUpdateGaeConfig(*g.Config)
}

// 打印所有服务提供者
func (g *GaeProviderManager) PrintProviders() {
	println("All of providers")
	for k, v := range g.Config.Providers {
		println(k + " -> " + v.Host)
	}
}

func ArgProvider(args []string) {
	m := NewGaeProviderManager()
	switch args[0] {
	case "add":
		m.AddProvider(args[1], args[2], args[3])
	case "remove":
		m.RemoveProvider(args[1])
	case "update":
		lowArg1 := strings.ToLower(args[1])
		switch lowArg1 {
		case "host":
			m.UpdateProviderHost(args[2], args[3])
		case "authkey":
			m.UpdateProviderAuthKey(args[2], args[3])
		default:
			println("update command not found")
		}
	case "default":
		m.DefaultProvider(args[1])
	default:
		m.PrintProviders()
	}
}
