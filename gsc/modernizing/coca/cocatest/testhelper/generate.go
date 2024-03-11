package testhelper

import (
	"gae-cli/gsc/modernizing/coca/pkg/application/analysis/javaapp"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"path/filepath"
)

func BuildAnalysisDeps(codePath string) ([]core_domain.CodeDataStruct, map[string]core_domain.CodeDataStruct, []core_domain.CodeDataStruct) {
	codePath = filepath.FromSlash(codePath)

	identifierApp := javaapp.NewJavaIdentifierApp()
	identifiers := identifierApp.AnalysisPath(codePath)

	callApp := javaapp.NewJavaFullApp()
	callNodes := callApp.AnalysisPath(codePath, identifiers)

	identifiersMap := core_domain.BuildIdentifierMap(identifiers)
	return callNodes, identifiersMap, identifiers
}
