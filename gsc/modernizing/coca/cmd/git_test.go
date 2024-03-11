package cmd

import (
	"gae-cli/gsc/modernizing/coca/cocatest/testcase"
	"testing"
)

// Todo: fake it
func TestGit(t *testing.T) {
	tests := []testcase.CmdTestCase{{
		Name:   "git",
		Cmd:    "git -a -f -t -b -o -r com -s 10 -m",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
