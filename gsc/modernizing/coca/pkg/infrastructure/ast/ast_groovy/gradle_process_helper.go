package ast_groovy

import (
	"gae-cli/gsc/modernizing/coca/languages/groovy"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func ProcessGroovyString(code string) *parser.GroovyParser {
	is := antlr.NewInputStream(code)
	lexer := parser.NewGroovyLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewGroovyParser(stream)
	return parser
}
