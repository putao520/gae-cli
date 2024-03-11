package cmd

import (
	"gae-cli/gsc/modernizing/coca/cocatest/testcase"
	"testing"
)

func Test_Analysis_Java(t *testing.T) {
	path := "config"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + path,
		Golden: "testdata/analysis_java.txt",
	}}
	RunTestCmd(t, analysis)
}
