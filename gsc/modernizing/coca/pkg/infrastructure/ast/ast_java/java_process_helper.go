package ast_java

import (
	"gae-cli/gsc/modernizing/coca/languages/java"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func ProcessJavaFile(path string) *parser.JavaParser {
	is, _ := antlr.NewFileStream(path)
	return processStream(is)
}

func processStream(is antlr.CharStream) *parser.JavaParser {
	lexer := parser.NewJavaLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	parser := parser.NewJavaParser(stream)
	return parser
}

func ProcessJavaString(code string) *parser.JavaParser {
	is := antlr.NewInputStream(code)
	return processStream(is)
}
