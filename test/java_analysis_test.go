package test

import (
	"gae-cli/handle/api_client"
	"os"
	"testing"
)

func TestAnalysisJava(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		panic("get current path error")
	}

	p := api_client.NewApiDescBuilderFromJava(dir)

	json, err := p.BuildApiJson("src/main/java/Api")
	if err != nil {
		println("build api json error")
	}
	println(json)
}

func TestUploaderApiDesc(t *testing.T) {
	err := api_client.UploaderApiDefinition()
	if err != nil {
		println("upload api definition error")
	}
}
