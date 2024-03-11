package cmd

import (
	"gae-cli/gsc/modernizing/coca/cocatest"
	"gae-cli/gsc/modernizing/coca/cocatest/testcase"
	"testing"
)

func RunTestCmd(t *testing.T, tests []testcase.CmdTestCase) {
	cocatest.RunTestCaseWithCmd(t, tests, NewRootCmd)
}
