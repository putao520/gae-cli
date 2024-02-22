package services

import (
	"gae-cli/gsc"
	"gae-cli/types"
)

type DeployItem struct {
	AppID           string         `json:"appid"`
	AppName         string         `json:"appname"`
	ClusterAddr     string         `json:"clusteraddr"`
	Config          string         `json:"config"`
	Container       int            `json:"container"`
	CreateAt        types.DateTime `json:"createat"`
	DataModel       string         `json:"datamodel"`
	Debug           int            `json:"debug"`
	ID              int            `json:"id"`
	Kind            string         `json:"kind"`
	Name            string         `json:"name"`
	Open            int            `json:"open"`
	ProxyTarget     string         `json:"proxy_target"`
	Replicaset      int            `json:"replicaset"`
	Secure          int            `json:"secure"`
	ServiceCategory string         `json:"service_category"`
	ServiceID       int            `json:"serviceid"`
	ServiceName     string         `json:"servicename"`
	State           int            `json:"state"`
	SubAddr         string         `json:"subaddr"`
	TargetPort      int            `json:"target_port"`
	Text            string         `json:"text"`
	UpdateAt        types.DateTime `json:"updateat"`
	UserID          string         `json:"userid"`
	Version         string         `json:"version"`
}

func QueryDeployInfo(appId string, serviceId string) *DeployItem {
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "ServicesDeploy").Call("findEx", []string{"ja:[{\"field\":\"appid\",\"value\":\"" + appId + "\",\"logic\":\"==\"},{\"field\":\"serviceid\",\"value\":\"" + serviceId + "\",\"logic\":\"==\"}]"})
	if !res.GetStatus() {
		println(res.GetMessage())
		return nil
	}
	result := res.AsObject()
	if result == nil {
		return nil
	}
	return gsc.MapToStruct[DeployItem](result)
}
