package evaluator

import (
	"fmt"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
)

type Controller struct {
}

func (Controller) Evaluate(node core_domain.CodeDataStruct) {
	fmt.Println("controller")
}
