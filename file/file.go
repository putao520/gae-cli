package file

import (
	"bufio"
	"fmt"
	"gae-cli/arg"
	"io"
	"os"
)

// 处理要使用的文件
func FileReplace(path string, env map[string]string) error {
	newPath := path + ".backup"

	if f, err := PathExists(path); f == false {
		println("." + path + " file not found")
		return err
	}

	// 备份文件已存在,删除当前文件,备份文件恢复成原始文件
	FileRestore(path)

	// 复制一份备份文件
	CopyFile(newPath, path)

	// 读备份文件
	var backupFile, bFileErr = os.OpenFile(newPath, os.O_RDONLY, 0666)
	defer backupFile.Close()
	if bFileErr != nil {
		println("." + newPath + " file can't open")
		return bFileErr
	}
	// 写原始文件
	var orgFile, oFileErr = os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0666)
	defer orgFile.Close()
	if oFileErr != nil {
		println("." + path + " file can't open")
		return oFileErr
	}

	writer := bufio.NewWriter(orgFile)
	scanner := bufio.NewScanner(backupFile)
	for scanner.Scan() {
		line := scanner.Text()
		newline := arg.EnvReplace(line, env)
		writer.Write([]byte(newline + "\n"))
	}
	writer.Flush()

	return nil
}

// 还原文件
func FileRestore(path string) {
	newPath := path + ".backup"
	if f, _ := PathExists(newPath); f == false {
		return
	}
	os.Remove(path)
	os.Rename(newPath, path)
}

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)

	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	defer srcFile.Close()

	//通过srcFile，获取到Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err = %v\n", err)
		return
	}

	writer := bufio.NewWriter(dstFile)
	defer func() {
		writer.Flush() //把缓冲区的内容写入到文件
		dstFile.Close()

	}()

	return io.Copy(writer, reader)

}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
