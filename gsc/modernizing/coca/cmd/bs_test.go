package cmd

import (
	"gae-cli/gsc/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestBadSmell(t *testing.T) {
	abs := "../_fixtures/bs"

	tests := []testcase.CmdTestCase{{
		Name:   "bs",
		Cmd:    "bs -s type -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, tests)
}
