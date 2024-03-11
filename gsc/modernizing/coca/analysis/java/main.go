package main

import (
	"gae-cli/gsc/modernizing/coca/analysis/java/app"
	"os"
)

func main() {
	output := os.Stdout
	rootCmd := app.NewRootCmd(output)
	_ = rootCmd.Execute()
}
