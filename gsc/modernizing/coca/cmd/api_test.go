package cmd

import (
	"gae-cli/gsc/modernizing/coca/cocatest/testcase"
	"testing"
)

func TestApi(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + path,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "api",
		Cmd:    "api -c -f -p " + path,
		Golden: "testdata/api.txt",
	}}
	RunTestCmd(t, tests)
}

func Test_ApiWithSortRemove(t *testing.T) {
	path := "../_fixtures/call"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + path,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "api",
		Cmd:    "api -c -s -r com.phodal.pholedge.book. -p" + path,
		Golden: "testdata/api_sort_remove.txt",
	}}
	RunTestCmd(t, tests)
}
