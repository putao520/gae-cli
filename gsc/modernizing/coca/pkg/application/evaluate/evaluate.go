package evaluate

import (
	"gae-cli/gsc/modernizing/coca/pkg/application/evaluate/evaluator"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
)

type Evaluator interface {
	Evaluate(result *evaluator.EvaluateModel, node core_domain.CodeDataStruct)
	EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []core_domain.CodeDataStruct, nodeMap map[string]core_domain.CodeDataStruct, identifiers []core_domain.CodeDataStruct)
}

type Evaluation struct {
	Evaluator Evaluator
}

func (o *Evaluation) Evaluate(result *evaluator.EvaluateModel, node core_domain.CodeDataStruct) {
	o.Evaluator.Evaluate(result, node)
}

func (o *Evaluation) EvaluateList(evaluateModel *evaluator.EvaluateModel, nodes []core_domain.CodeDataStruct, nodeMap map[string]core_domain.CodeDataStruct, identifiers []core_domain.CodeDataStruct) {
	o.Evaluator.EvaluateList(evaluateModel, nodes, nodeMap, identifiers)
}
