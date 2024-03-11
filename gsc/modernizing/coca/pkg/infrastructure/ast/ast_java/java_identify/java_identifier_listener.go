package java_identify

import (
	"gae-cli/gsc/modernizing/coca/languages/java"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java/common_listener"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"reflect"
	"strings"
)

var currentNode *core_domain.CodeDataStruct
var nodes []core_domain.CodeDataStruct

var currentMethod core_domain.CodeFunction
var hasEnterClass = false
var imports []string

func NewJavaIdentifierListener() *JavaIdentifierListener {
	nodes = nil
	currentNode = core_domain.NewDataStruct()
	currentMethod = core_domain.NewJMethod()
	return &JavaIdentifierListener{}
}

type JavaIdentifierListener struct {
	parser.BaseJavaParserListener
}

func (s *JavaIdentifierListener) EnterImportDeclaration(ctx *parser.ImportDeclarationContext) {
	importText := ctx.QualifiedName().GetText()
	imports = append(imports, importText)
}

func (s *JavaIdentifierListener) EnterPackageDeclaration(ctx *parser.PackageDeclarationContext) {
	currentNode.Package = ctx.QualifiedName().GetText()
}

func (s *JavaIdentifierListener) EnterClassDeclaration(ctx *parser.ClassDeclarationContext) {
	hasEnterClass = true

	currentNode.Type = "Class"
	if ctx.Identifier() != nil {
		currentNode.NodeName = ctx.Identifier().GetText()
	}

	if ctx.EXTENDS() != nil {
		if ctx.TypeType() != nil {
			currentNode.Extend = ctx.TypeType().GetText()
		}
	}

	if ctx.IMPLEMENTS() != nil {
		list := ctx.TypeList(0)
		switch x := list.(type) {
		case *parser.TypeListContext:
			{
				types := x.AllTypeType()
				for _, typ := range types {
					typeText := typ.GetText()
					for _, imp := range imports {
						if strings.HasSuffix(imp, "."+typeText) {
							currentNode.Implements = append(currentNode.Implements, imp)
						}
					}
				}
			}
		}

	}

	currentMethod = core_domain.NewJMethod()
}

func (s *JavaIdentifierListener) ExitClassBody(ctx *parser.ClassBodyContext) {
	hasEnterClass = false
	if currentNode.NodeName != "" {
		nodes = append(nodes, *currentNode)
	}
	currentNode = core_domain.NewDataStruct()
}

func (s *JavaIdentifierListener) EnterConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	position := core_domain.CodePosition{
		StartLine:         ctx.GetStart().GetLine(),
		StartLinePosition: ctx.GetStart().GetColumn(),
		StopLine:          ctx.GetStop().GetLine(),
		StopLinePosition:  ctx.GetStop().GetColumn(),
	}

	currentMethod = core_domain.CodeFunction{
		Name:          ctx.Identifier().GetText(),
		ReturnType:    "",
		Override:      isOverrideMethod,
		Annotations:   currentMethod.Annotations,
		IsConstructor: true,
		Position:      position,
	}
}

func (s *JavaIdentifierListener) ExitConstructorDeclaration(ctx *parser.ConstructorDeclarationContext) {
	currentNode.Functions = append(currentNode.Functions, currentMethod)
}

func buildMethodParameters(parameters parser.IFormalParametersContext, method *core_domain.CodeFunction) {
	if parameters != nil {
		if parameters.GetChild(0) == nil || parameters.GetText() == "()" || parameters.GetChild(1) == nil {
			return
		}
		method.Parameters = ast_java.BuildMethodParameters(parameters)
	}
}

var isOverrideMethod = false

func (s *JavaIdentifierListener) EnterMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	hasEnterClass = true

	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ""
	if ctx.Identifier() != nil {
		name = ctx.Identifier().GetText()
	}

	typeType := ctx.TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	position := core_domain.CodePosition{
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
	}

	currentMethod = core_domain.CodeFunction{
		Name:        name,
		ReturnType:  typeType,
		Override:    isOverrideMethod,
		Annotations: currentMethod.Annotations,
		Position:    position,
	}

	if reflect.TypeOf(ctx.GetParent().GetParent()).String() == "*parser.ClassBodyDeclarationContext" {
		bodyCtx := ctx.GetParent().GetParent().(*parser.ClassBodyDeclarationContext)
		for _, modifier := range bodyCtx.AllModifier() {
			if !strings.Contains(modifier.GetText(), "@") {
				currentMethod.Modifiers = append(currentMethod.Modifiers, modifier.GetText())
			}
		}
	}

	parameters := ctx.FormalParameters()
	buildMethodParameters(parameters, &currentMethod)

	isOverrideMethod = false
}

func (s *JavaIdentifierListener) ExitMethodDeclaration(ctx *parser.MethodDeclarationContext) {
	currentNode.Functions = append(currentNode.Functions, currentMethod)
	currentMethod = core_domain.NewJMethod()
}

func (s *JavaIdentifierListener) EnterAnnotation(ctx *parser.AnnotationContext) {
	if ctx.QualifiedName() == nil {
		return
	}

	annotationName := ctx.QualifiedName().GetText()
	if annotationName == "Override" {
		isOverrideMethod = true
	}

	if !hasEnterClass {
		annotation := common_listener.BuildAnnotation(ctx)
		currentNode.Annotations = append(currentNode.Annotations, annotation)
	}
}

func (s *JavaIdentifierListener) EnterInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = true
	currentNode.Type = "Interface"
	currentNode.NodeName = ctx.Identifier().GetText()
}

func (s *JavaIdentifierListener) ExitInterfaceDeclaration(ctx *parser.InterfaceDeclarationContext) {
	hasEnterClass = false
	if currentNode.NodeName != "" {
		nodes = append(nodes, *currentNode)
	}
	currentNode = core_domain.NewDataStruct()
}

func (s *JavaIdentifierListener) EnterInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	startLine := ctx.GetStart().GetLine()
	startLinePosition := ctx.GetStart().GetColumn()
	stopLine := ctx.GetStop().GetLine()
	stopLinePosition := ctx.GetStop().GetColumn()
	name := ctx.InterfaceCommonBodyDeclaration().(*parser.InterfaceCommonBodyDeclarationContext).Identifier().GetText()
	//XXX: find the start position of {, not public
	typeType := ctx.InterfaceCommonBodyDeclaration().(*parser.InterfaceCommonBodyDeclarationContext).TypeTypeOrVoid().GetText()

	if reflect.TypeOf(ctx.GetParent().GetParent().GetChild(0)).String() == "*parser.ModifierContext" {
		common_listener.BuildAnnotationForMethod(ctx.GetParent().GetParent().GetChild(0).(*parser.ModifierContext), &currentMethod)
	}

	position := core_domain.CodePosition{
		StartLine:         startLine,
		StartLinePosition: startLinePosition,
		StopLine:          stopLine,
		StopLinePosition:  stopLinePosition,
	}

	currentMethod = core_domain.CodeFunction{
		Name:        name,
		ReturnType:  typeType,
		Override:    isOverrideMethod,
		Annotations: currentMethod.Annotations,
		Position:    position,
	}
}

func (s *JavaIdentifierListener) ExitInterfaceMethodDeclaration(ctx *parser.InterfaceMethodDeclarationContext) {
	currentNode.Functions = append(currentNode.Functions, currentMethod)
	currentMethod = core_domain.NewJMethod()
}

func (s *JavaIdentifierListener) EnterExpression(ctx *parser.ExpressionContext) {
	if reflect.TypeOf(ctx.GetParent()).String() == "*parser.StatementContext" {
		statementCtx := ctx.GetParent().(*parser.StatementContext)
		firstChild := statementCtx.GetChild(0).(antlr.ParseTree).GetText()
		if strings.ToLower(firstChild) == "return" {
			currentMethod.IsReturnNull = strings.Contains(ctx.GetText(), "null")
		}
	}
}

func (s *JavaIdentifierListener) GetNodes() []core_domain.CodeDataStruct {
	return nodes
}

func (s *JavaIdentifierListener) GetImports() []string {
	return imports
}
