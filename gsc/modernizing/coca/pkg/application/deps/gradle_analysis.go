package deps

import (
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_groovy"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func AnalysisGradleFile(path string) []core_domain.CodeDependency {
	bytes := cmd_util.ReadFile(path)
	return AnalysisGradleString(string(bytes))
}

func AnalysisGradleString(str string) []core_domain.CodeDependency {
	parser := ast_groovy.ProcessGroovyString(str)
	context := parser.CompilationUnit()
	listener := ast_groovy.NewGroovyIdentListener()
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetDepsInfo()
}
