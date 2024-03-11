package api_client

import (
	"errors"
	"github.com/beevik/etree"
	"github.com/evilsocket/islazy/zip"
	"os"
)

var GlobalTmpDir string = ""

// 获得 maven 本地仓库路径
func LoadLocalMavenRepository() (string, error) {
	// 读取环境变量 MAVEN_HOME 的值
	mavenHome := os.Getenv("MAVEN_HOME")
	if mavenHome == "" {
		println("环境变量 MAVEN_HOME 未设置")
		return "", errors.New("环境变量 MAVEN_HOME 未设置")
	}
	// 读取 maven 本地仓库路径
	mavenConfigPath := mavenHome + "/conf/settings.xml"
	mavenConfigContent, err := os.ReadFile(mavenConfigPath)
	if err != nil {
		println("读取 maven 配置文件失败")
		return "", err
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(mavenConfigContent); err != nil {
		panic(err)
	}
	// 读取本地仓库路径
	localRepository := doc.FindElement("//localRepository")
	if localRepository == nil {
		println("未找到本地仓库路径")
		return "", errors.New("未找到本地仓库路径")
	}
	return localRepository.Text(), nil
}

// 构造解析 maven 依赖的结构体, 返回临时文件夹路径
func LoadPomDependencyPackage() (string, error) {
	if GlobalTmpDir != "" {
		return GlobalTmpDir, nil
	}

	localRepository, err := LoadLocalMavenRepository()
	if err != nil {
		println("load local maven repository error")
		return "", err
	}
	dependencies, err := LoadPomDependencies()
	if err != nil {
		println("load pom dependencies error")
		return "", err
	}
	// 创建临时文件夹
	tmpDir := os.TempDir() + "/gae-cli"
	err = os.MkdirAll(tmpDir, os.ModePerm)
	if err != nil {
		println("create tmp dir error")
		return "", err
	}
	// 下载依赖包
	for _, dependency := range dependencies {
		// 构建依赖包路径
		p := "/" + Package2Path(dependency.GroupId) + "/" + dependency.ArtifactId + "/" + dependency.Version + "/" + dependency.ArtifactId + "-" + dependency.Version + "." + dependency.Packaging
		dependencyPath := localRepository + p
		// 解压缩依赖包
		_, err = zip.Unzip(dependencyPath, tmpDir)
		if err != nil {
			println(dependencyPath + " -> unzip dependency error")
			_ = os.RemoveAll(tmpDir)
			return "", err
		}
	}
	GlobalTmpDir = tmpDir
	return tmpDir, nil
}
