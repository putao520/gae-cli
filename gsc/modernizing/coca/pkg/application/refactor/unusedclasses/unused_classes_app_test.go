package unusedclasses

import (
	"encoding/json"
	"gae-cli/gsc/modernizing/coca/cmd/cmd_util"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	. "github.com/onsi/gomega"
	"path/filepath"
	"testing"
)

func TestRefactoring(t *testing.T) {
	g := NewGomegaWithT(t)

	var parsedDeps []core_domain.CodeDataStruct
	codePath := "../../../../_fixtures/count/call.json"
	codePath = filepath.FromSlash(codePath)
	file := cmd_util.ReadFile(codePath)
	_ = json.Unmarshal(file, &parsedDeps)

	results := Refactoring(parsedDeps)

	g.Expect(len(results)).To(Equal(1))
}
