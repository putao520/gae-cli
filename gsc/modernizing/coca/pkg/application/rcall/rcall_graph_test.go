package rcall

import (
	"encoding/json"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"log"
	"testing"

	. "github.com/onsi/gomega"
)

func MockWriteCallMap(rcallMap map[string][]string) {

}

func TestRCallGraph_Analysis(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewRCallGraph()
	file := cmd_util.ReadFile("../../../_fixtures/call/call_api_test.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	content := analyser.Analysis("com.phodal.pholedge.book.BookService.createBook", parsedDeps, MockWriteCallMap)

	g.Expect(content).To(Equal(`digraph G {
"com.phodal.pholedge.book.BookController.createBook" -> "com.phodal.pholedge.book.BookService.createBook";
}
`))
}

func TestRCallGraph_Constructor(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	analyser := NewRCallGraph()
	file := cmd_util.ReadFile("../../../_fixtures/rcall/constructor_call.json")
	if file == nil {
		log.Fatal("lost file")
	}

	_ = json.Unmarshal(file, &parsedDeps)

	content := analyser.Analysis("com.phodal.coca.analysis.JavaCallApp.parse", parsedDeps, MockWriteCallMap)

	// Todo bug: to be fix
	g.Expect(content).To(Equal(`digraph G {
"com.phodal.coca.analysis.JavaCallApp.analysisDir" -> "com.phodal.coca.analysis.JavaCallApp.parse";
}
`))
}
