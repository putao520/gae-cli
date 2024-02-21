package handle

import (
	"bufio"
	"fmt"
	"gae-cli/arg"
	"gae-cli/file"
	"os"
	"os/exec"
	"syscall"
)

func ArgRun(args []string) {
	// 构造入参 hashmap
	env := arg.ChildArgBuild(args)
	// 替换 hashmap 内变量
	for key, val := range env {
		env[key] = arg.EnvReplace(val, env)
		// println(env[key])
	}

	path, err := os.Getwd()
	if err != nil {
		return
	}

	// 处理 docker 文件
	file.FileReplace("dockerfile", env)
	defer file.FileRestore("dockerfile")

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

func doShell(cmdString string, env map[string]string) (string, error) {
	// 替换参数
	cmdString = arg.EnvReplace(cmdString, env)

	// 分离参数与命令
	cmdArr := arg.CommandlineToArray(cmdString)
	command := cmdArr[0]
	args := cmdArr[1:]

	// 执行内置命令
	if next := arg.CmdFilter(command, args...); !next {
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
