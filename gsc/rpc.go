package gsc

import "gae-cli/config"

var rpcMasterClient *RpcRequest
var rpcClient *RpcRequest

func GetRpcClient() *RpcRequest {
	if rpcClient == nil {
		gaeManager := config.NewGaeConfigManager()

		p := gaeManager.GetDefaultProvider()
		if p == nil {
			println("GetDefaultProvider is invalid")
			return nil
		}

		var appId string = "0"
		projectCfg := config.ReadGaeProjectConfig()
		if projectCfg != nil {
			appId = projectCfg.Project.AppID
		}

		rpcClient = NewRpcClient(appId, p.AuthKey, p.Host, p.Token)
	}
	return rpcClient
}

func GetRpcMasterClient() *RpcRequest {
	if rpcMasterClient == nil {
		gaeManager := config.NewGaeConfigManager()

		p := gaeManager.GetDefaultProvider()
		if p == nil {
			panic("GetDefaultProvider is invalid")
		}

		rpcMasterClient = NewRpcMasterClient(p.AuthKey, p.Host, p.Token)
	}
	return rpcMasterClient
}
