package test

import (
	"gae-cli/handle/api_client"
	"testing"
)

func TestMaven(t *testing.T) {
	localRepository, err := api_client.LoadLocalMavenRepository()
	if err != nil {
		println("load local maven repository error")
	}
	println(localRepository)
}

func TestPom(t *testing.T) {
	dependencies, err := api_client.LoadPomDependencies()
	if err != nil {
		println("load pom dependencies error")
	}
	for _, dependency := range dependencies {
		println(dependency.GroupId)
		println(dependency.ArtifactId)
		println(dependency.Version)
		println(dependency.Packaging)
	}
}

func TestMavenPackageFolderBuilder(t *testing.T) {
	tmpDir, err := api_client.LoadPomDependencyPackage()
	if err != nil {
		println("load pom dependencies package error")
	}
	defer func() {
		// 删除临时文件夹
		if err == nil {
			// err := os.RemoveAll(tmpDir)
			if err != nil {
				return
			}
		}
	}()

	p := api_client.NewApiDescBuilderFromClass(tmpDir)
	desc, err := p.BuildApiJson("common/java/ServiceTemplate/MicroServiceTemplate.class")
	if err != nil {
		println("build api json error")
	}
	println(desc.Name)

	println(tmpDir)
}
