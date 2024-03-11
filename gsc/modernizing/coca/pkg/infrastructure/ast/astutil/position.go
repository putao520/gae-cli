package astutil

import (
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func AddFunctionPosition(m *core_domain.CodeFunction, ctx *antlr.BaseParserRuleContext) {
	m.Position.StartLine = ctx.GetStart().GetLine()
	m.Position.StartLinePosition = ctx.GetStart().GetColumn()
	m.Position.StopLine = ctx.GetStop().GetLine()
	m.Position.StopLinePosition = ctx.GetStop().GetColumn()
}
