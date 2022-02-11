package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"syscall"
)

var globalVer = "0.0.1"

/**
 	version
	run -key value,...
	help
*/
func main() {
	// 访问当前上下文的目录，获得 .gpaeshell 文件，依次支持它
	println("|--------------------------|")
	println("|     Grape App Engine     |")
	println("|               v0.0.1     |")
	println("|--------------------------|")

	argFilter(os.Args)
}

func argFilter(args []string) {
	if len(args) == 1 {
		return
	}
	cmd := args[1]
	switch strings.ToLower(cmd) {
	case "version", "v":
		argVersion()
	case "run", "r":
		argRun(args[2:])
	case "help", "h":
		argHelp()
	default:
		unknownArg(cmd)
	}
}

func unknownArg(arg string) {
	println("GAE: '" + arg + "' is not a GAE command.\n")
	println("See gae help\n")
}

func argHelp() {
	println("Commands")
	fmt.Printf("\t%-20s %s\n", "run, r", "Run GAE's script")
	fmt.Printf("\t%-20s %s\n", "version, v", "Show the GAE version information")
	fmt.Printf("\t%-20s %s\n", "help, h", "Show the GAE help invformation")
}

func argVersion() {
	fmt.Printf("GAE Version:%s", globalVer)
}

func childArgBuild(args []string) map[string]string {
	result := make(map[string]string)
	for idx, arg := range args {
		if strings.HasPrefix(arg, "-") && (idx+1) < len(args) {
			result[arg] = args[idx+1]
		}
		// println(arg)
	}
	return result
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

// 处理要使用的文件
func fileReplace(path string, env map[string]string) error {
	newPath := path + ".backup"

	if f, err := PathExists(path); f == false {
		println("." + path + " file not found")
		return err
	}

	// 备份文件已存在,删除当前文件,备份文件恢复成原始文件
	fileRestore(path)

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
		newline := envReplace(line, env)
		writer.Write([]byte(newline + "\n"))
	}
	writer.Flush()

	return nil
}

// 还原文件
func fileRestore(path string) {
	newPath := path + ".backup"
	if f, _ := PathExists(newPath); f == false {
		return
	}
	os.Remove(path)
	os.Rename(newPath, path)
}

func argRun(args []string) {
	// 构造入参 hashmap
	env := childArgBuild(args)
	// 替换 hashmap 内变量
	for key, val := range env {
		env[key] = envReplace(val, env)
	}

	path, err := os.Getwd()
	if err != nil {
		return
	}

	// 处理 docker 文件
	fileReplace("dockerfile", env)
	defer fileRestore("dockerfile")

	// 处理 gaeshell 文件
	shellPath := path + "/.gaeshell"
	if _, err := os.Stat(shellPath); os.IsExist(err) {
		println(".gaeshell file not found")
		return
	}

	var shell, fileErr = os.OpenFile(shellPath, os.O_RDONLY, 0666)
	if fileErr != nil {
		println(".gaeshell file can't open")
		return
	}
	defer shell.Close()

	scanner := bufio.NewScanner(shell)
	for scanner.Scan() {
		line := scanner.Text()
		if cmdString, err := doShell(line, env); err != nil {
			fmt.Printf("Command:%s ->Error:%s\n", cmdString, err)
			return // 如果脚本执行出错，禁止脚本运行
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Cannot scanner text file: %s, err: [%v]\n", shellPath, err)
		return
	}
}

func testShell() {
	cmd := exec.Command("docker", "version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("message: %s\n", output)
}

func doShell(cmdString string, env map[string]string) (string, error) {
	// 替换参数
	cmdString = envReplace(cmdString, env)

	// 分离参数与命令
	cmdArr := commandlineToArray(cmdString)
	command := cmdArr[0]
	args := cmdArr[1:]

	// 执行内置命令
	if next := cmdFilter(command, args...); !next {
		return cmdString, nil
	}

	// 执行系统命令
	cmd := exec.Command(command)
	// 还原命令行
	cmd.SysProcAttr = &syscall.SysProcAttr{CmdLine: cmdString}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	/*
		output, err := cmd.CombinedOutput()
		if err != nil {
			return "", err
		}
	*/

	// 执行时，包含双引号的参数被转义成 "\""
	if err := cmd.Run(); err != nil {
		return cmdString, err
	}

	return cmdString, nil
}

func commandlineToArray(cmdString string) []string {
	result := make([]string, 0)
	spaceStart := 0
	quotesMode := false
	l := len(cmdString)
	// 按双引号和空格处理数组
	for i := 0; i < l; i++ {
		char := cmdString[i]
		// 按空格分组
		if char == ' ' && quotesMode == false {
			if i > 0 {
				result = append(result, cmdString[spaceStart:i])
				spaceStart = i + 1
			}
		}
		// 遇到双引号时，以双引号为分割块(转移符双引号不算双引号)
		if char == '"' && i > 0 && cmdString[i-1] != '\\' {
			if quotesMode == false {
				spaceStart = i
			}
			quotesMode = !quotesMode
		}
	}
	// 处理尾部截止数值
	if l > spaceStart {
		result = append(result, cmdString[spaceStart:])
	}
	return result
}

// 替换命令行内变量
func envReplace(command string, env map[string]string) string {
	startIdx := 0
	inlineMode := false
	dataDict := make([]string, 0)
	l := len(command)
	for i := 0; i < l; i++ {
		// 找到环境变量入口
		if command[i] == '#' && (i+1) < l && command[i+1] == '{' {
			dataDict = append(dataDict, command[startIdx:i])
			startIdx = i + 2
			inlineMode = true
		}
		// 找到环境变量闭合
		if command[i] == '}' && inlineMode == true {
			dataDict = append(dataDict, "-"+command[startIdx:i])
			inlineMode = false
			startIdx = i + 1
		}
	}
	if l > startIdx {
		dataDict = append(dataDict, command[startIdx:])
	}
	// 替换值
	for idx, val := range dataDict {
		if strings.HasPrefix(val, "-") {
			v, f := env[val]
			if f {
				dataDict[idx] = v
			}
		}
	}
	// 还原
	return strings.Join(dataDict, "")
}

func cmdFilter(command string, arg ...string) bool {
	switch command {
	case "echo":
		cmdEcho(arg[0])
		return false
	}
	return true
}

func cmdEcho(arg string) {
	println(strings.Trim(arg, "\""))
}
