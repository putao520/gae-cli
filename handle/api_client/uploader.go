package api_client

import (
	"encoding/json"
	"fmt"
	"gae-cli/config"
	"gae-cli/gsc"
	"os"
)

func UploaderApiDefinition() error {
	dir, err := os.Getwd()
	if err != nil {
		panic("get current path error")
	}

	apiJava := NewApiDescBuilderFromJava(dir)

	apiDesc, err := apiJava.BuildApiJson("src/main/java/Api")
	if err != nil {
		fmt.Printf("构建api描述时出错: %v\n", err)
		return err
	}

	// 读取当前项目配置
	projectConfig := config.ReadGaeProjectConfig()
	serviceId := projectConfig.Project.ServiceID
	// 构造上传数据
	inputData := make(map[string]interface{})
	// apiDesc 转成 json string
	apiDescStr, err := json.Marshal(apiDesc)
	if err != nil {
		fmt.Printf("apiDesc转json时出错: %v\n", err)
		return err
	}
	inputData["api"] = string(apiDescStr)
	// inputData 转成 json string
	inputDataStr, err := json.Marshal(inputData)
	if err != nil {
		fmt.Printf("inputData转json时出错: %v\n", err)
		return err
	}
	paramArrayString := "j:" + string(inputDataStr)
	res := gsc.GetRpcMasterClient().SetEndpoint("system", "Service").Call("update", []string{serviceId, paramArrayString})
	if !res.GetStatus() {
		fmt.Printf("上传api描述时出错: %v\n", res.GetMessage())
		return fmt.Errorf(res.GetMessage())
	}
	return nil
}
