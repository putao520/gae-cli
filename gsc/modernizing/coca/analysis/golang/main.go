package main

import (
	"gae-cli/gsc/modernizing/coca/analysis/golang/app"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := app.NewRootCmd(output)
	_ = rootCmd.Execute()
}
