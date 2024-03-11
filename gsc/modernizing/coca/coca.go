package main

import (
	"fmt"
	"gae-cli/gsc/modernizing/coca/cmd"
	"github.com/pkg/profile"
	"os"
	"time"
)

func main() {
	t1 := time.Now() // get current time
	defer profile.Start().Stop()
	output := os.Stdout
	rootCmd := cmd.NewRootCmd(output)
	_ = rootCmd.Execute()
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
