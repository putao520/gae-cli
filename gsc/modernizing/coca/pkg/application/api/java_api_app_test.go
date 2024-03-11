package api

import (
	"gae-cli/gsc/modernizing/coca/cocatest/testhelper"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	. "github.com/onsi/gomega"
	"testing"
)

func TestJavaCallApp_AnalysisPath(t *testing.T) {
	g := NewGomegaWithT(t)

	codePath := "../../../_fixtures/call"
	callNodes, identifiersMap, identifiers := testhelper.BuildAnalysisDeps(codePath)
	diMap := core_domain.BuildDIMap(identifiers, identifiersMap)

	app := new(JavaApiApp)
	restApis := app.AnalysisPath(codePath, callNodes, identifiersMap, diMap)

	g.Expect(len(restApis)).To(Equal(4))
	g.Expect(restApis[0].HttpMethod).To(Equal("POST"))
	g.Expect(restApis[0].Uri).To(Equal("/books"))
}
