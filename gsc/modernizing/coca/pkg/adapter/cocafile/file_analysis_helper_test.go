package cocafile

import (
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java/java_identify"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_ProcessStringWorks(t *testing.T) {
	g := NewGomegaWithT(t)
	parser := ast_java.ProcessJavaString(`
package com.phodal.coca.analysis.identifier.model;

public class DataClass {
    private String date;

    public String getDate() {
        return date;
    }
}

`)

	context := parser.CompilationUnit()
	listener := java_identify.NewJavaIdentifierListener()

	antlr.NewParseTreeWalker().Walk(listener, context)

	identifiers := listener.GetNodes()
	g.Expect(identifiers[0].NodeName).To(Equal("DataClass"))
}
