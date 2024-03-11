package analysis

import (
	"encoding/json"
	"fmt"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/pkg/adapter/cocafile"
	"gae-cli/gsc/modernizing/coca/pkg/application/analysis/app_concept"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"io"
	"io/ioutil"
	"unicode"
)

func CommonAnalysis(output io.Writer, path string, app app_concept.AbstractAnalysisApp, filter func(path string) bool, isFunctionBase bool) []core_domain.CodeDataStruct {
	var results []core_domain.CodeContainer
	files := cocafile.GetFilesWithFilter(path, filter)

	var codeMembers []core_domain.CodeMember

	app.AnalysisPackageManager(path)

	for _, file := range files {
		content, _ := ioutil.ReadFile(file)
		members := app.IdentAnalysis(string(content), file)
		codeMembers = append(codeMembers, members...)

		identModel, _ := json.MarshalIndent(codeMembers, "", "\t")
		cmd_util.WriteToCocaFile("members.json", string(identModel))
	}

	for _, file := range files {
		fmt.Fprintf(output, "Process file: %s\n", file)
		content, _ := ioutil.ReadFile(file)
		app.SetExtensions(codeMembers)
		result := app.Analysis(string(content), file)
		results = append(results, result)
	}

	var ds []core_domain.CodeDataStruct
	for _, result := range results {
		ds = append(ds, result.DataStructures...)

		if isFunctionBase {
			methodDs := BuildMethodDs(result)
			ds = append(ds, methodDs...)
		}
	}

	return ds
}

func BuildMethodDs(result core_domain.CodeContainer) []core_domain.CodeDataStruct {
	var methodsDs []core_domain.CodeDataStruct
	for _, member := range result.Members {
		for _, node := range member.FunctionNodes {
			if unicode.IsUpper(rune(node.Name[0])) {
				methodDs := core_domain.CodeDataStruct{
					NodeName:      node.Name,
					Package:       result.PackageName,
					FunctionCalls: node.FunctionCalls,
				}
				methodsDs = append(methodsDs, methodDs)
			}
		}
	}

	return methodsDs
}
