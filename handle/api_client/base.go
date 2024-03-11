package api_client

import (
	"fmt"
	"os"
	"path/filepath"
)

func IsGscServiceProject() bool {
	dir, err := os.Getwd()
	if err != nil {
		println("get current path error")
		return false
	}

	// 判断当前目录下是否包含 pom.xml 文件(无视大小写)
	// 构建完整的文件路径
	pomFilePath := filepath.Join(dir, "pom.xml")

	// 检查文件是否存在
	_, err = os.Stat(pomFilePath)
	if os.IsNotExist(err) {
		fmt.Println("当前目录不包含 pom.xml 文件")
		return false
	}

	// 判断当前目录下是否包含 gfw.cfg 文件(无视大小写)
	gfwFilePath := filepath.Join(dir, "gfw.cfg")
	_, err = os.Stat(gfwFilePath)
	if os.IsNotExist(err) {
		fmt.Println("当前目录不包含 gfw.cfg 文件")
		return false
	}
	return true
}

func GetFileName(path string) (string, string) {
	// 获取文件扩展名
	fileExt := filepath.Ext(path)

	// 去除扩展名后的文件名
	nameNoExt := path[:len(path)-len(fileExt)]

	return nameNoExt, fileExt
}
