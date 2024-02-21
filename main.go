package main

import (
	"gae-cli/arg"
	"gae-cli/gsc/services"
	"gae-cli/handle"
	"os"
	"strings"
)

/*
*

	 	version
		run -key value,...
		help
*/
func main() {
	// 访问当前上下文的目录，获得 .gpaeshell 文件，依次支持它
	println("|--------------------------|")
	println("|     Grape App Engine     |")
	println("|               v" + handle.GlobalVer + "     |")
	println("|--------------------------|")

	ArgFilter(os.Args)
}

func ArgFilter(args []string) {
	if len(args) == 1 {
		return
	}
	cmd := args[1]
	switch strings.ToLower(cmd) {
	case "version", "v":
		handle.ArgVersion()
	case "run", "r":
		handle.ArgRun(args[2:])
	case "provider":
		handle.ArgProvider(args[2:])
	case "login":
		services.Login(args[2], args[3])
	case "logout":
		services.Logout()
	case "help", "h":
		handle.ArgHelp()
	default:
		arg.UnknownArg(cmd)
	}
}
