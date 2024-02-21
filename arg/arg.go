package arg

import (
	"fmt"
	"os/exec"
	"strings"
)

// 替换命令行内变量
func EnvReplace(command string, env map[string]string) string {
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

func CmdFilter(command string, arg ...string) bool {
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

func ChildArgBuild(args []string) map[string]string {
	result := make(map[string]string)
	for idx, arg := range args {
		if strings.HasPrefix(arg, "-") && (idx+1) < len(args) {
			result[arg] = args[idx+1]
		}
	}
	return result
}

func testShell() {
	cmd := exec.Command("docker", "version")

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
	fmt.Printf("message: %s\n", output)
}

func CommandlineToArray(cmdString string) []string {
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

func UnknownArg(arg string) {
	println("GAE: '" + arg + "' is not a GAE command.\n")
	println("See gae help\n")
}
