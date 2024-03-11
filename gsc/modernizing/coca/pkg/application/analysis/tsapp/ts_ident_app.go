package tsapp

import (
	parser "gae-cli/gsc/modernizing/coca/languages/ts"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_typescript"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func processStream(is antlr.CharStream) *parser.TypeScriptParser {
	lexer := parser.NewTypeScriptLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	return parser.NewTypeScriptParser(stream)
}

func ProcessTsString(code string) *parser.TypeScriptParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}

type TypeScriptIdentApp struct {
}

func (t *TypeScriptIdentApp) AnalysisPackageManager(path string) core_domain.CodePackageInfo {
	return core_domain.CodePackageInfo{}
}

func (t *TypeScriptIdentApp) Analysis(code string, fileName string) core_domain.CodeContainer {
	scriptParser := ProcessTsString(code)
	context := scriptParser.Program()

	listener := ast_typescript.NewTypeScriptIdentListener(fileName)
	antlr.NewParseTreeWalker().Walk(listener, context)

	return listener.GetNodeInfo()
}

func (t *TypeScriptIdentApp) SetExtensions(extension interface{}) {

}

func (t *TypeScriptIdentApp) IdentAnalysis(s string, file string) []core_domain.CodeMember {
	return nil
}
