package api_client

import (
	"fmt"
	"gae-cli/gsc/modernizing/coca/pkg/domain/core_domain"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java"
	"gae-cli/gsc/modernizing/coca/pkg/infrastructure/ast/ast_java/java_identify"
	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"os"
	"path/filepath"
	"strings"
)

func buildSpec(annotations []core_domain.CodeAnnotation) []ApiSpec {
	// 遍历 annotations
	specs := make([]ApiSpec, 0)
	for _, annotation := range annotations {
		if annotation.Name == "InterfaceType" {
			v := annotation.KeyValues
			for _, kv := range v {
				vp := getSpecStr(kv.Value)
				if vp == CloseApi || vp == PrivateApi {
					return []ApiSpec{}
				}
				specs = append(specs, vp)
			}
		}
	}
	if len(specs) == 0 {
		specs = append(specs, PublicApi)
	}
	return specs
}

// 查找類引用
func findClassRef(imports []string, className string) string {
	for _, imp := range imports {
		if strings.HasSuffix(imp, "."+className) {
			return imp
		}
	}
	return ""
}

func (p *JavaApiDescBuilder) tryBuildApiDescription(path string) (ApiDescription, error) {
	// 優先本項目類查找
	apiDesc, err := p.buildApiDescription(path)
	if err != nil {
		// 如果err = 文件不存在，則去maven依賴中查找
		if os.IsNotExist(err) {
			path = strings.Join(strings.Split(path, "."), "/")
			if !strings.HasSuffix(path, ".class") {
				path = path + ".class"
			}
			apiDesc, err = p.mavenHelper.BuildApiJson(path)
		} else {
			fmt.Printf("build api desc [%s] error: %v\n", path, err)
			return ApiDescription{}, err
		}
	}
	return apiDesc, err
}

// 父類接口類型生成
func (p *JavaApiDescBuilder) fatherClassBuild(fatherClassName string, imports []string) ([]ApiAction, error) {
	actions := make([]ApiAction, 0)
	if fatherClassName != "" {
		// 查找父类引用
		imp := findClassRef(imports, fatherClassName)
		if imp == "" {
			println("fatherClassName: ", fatherClassName)
			return actions, fmt.Errorf("father class not found")
		}
		// 解析父类
		apiDesc, err := p.tryBuildApiDescription(imp)
		if err != nil {
			fmt.Printf("build father class [%s] error: %v\n", imp, err)
			return actions, err
		}
		actions = append(actions, apiDesc.Actions...)
	}
	return actions, nil
}

// 合并父类和子类的接口
func mixActions(parent []ApiAction, a ApiAction) []ApiAction {
	// 参数也一样
	same := true
	// 查找是否有重复的接口
	for _, p := range parent {
		if p.Name == a.Name && len(p.ParamsArray) == len(a.ParamsArray) {
			for i, param := range p.ParamsArray {
				if param.Type != a.ParamsArray[i].Type {
					same = false
					break
				}
			}
		}
	}
	if !same {
		parent = append(parent, a)
	}
	return parent
}

func (p *JavaApiDescBuilder) buildApiDescription(path string) (ApiDescription, error) {
	// 读取文件内容
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("读取文件内容时出错: %v\n", err)
		return ApiDescription{}, err
	}

	parserHandler := ast_java.ProcessJavaString(string(content))
	context := parserHandler.CompilationUnit()
	listener := java_identify.NewJavaIdentifierListener()
	antlr.NewParseTreeWalker().Walk(listener, context)
	identifiers := listener.GetNodes()
	// 只管第一个类
	clsInfo := identifiers[0]
	// 處理父類接口
	actions, err := p.fatherClassBuild(clsInfo.Extend, listener.GetImports())
	if err != nil {
		println("Parent class build error")
		return ApiDescription{}, err
	}

	functions := clsInfo.Functions
	// 遍历 functions
	for _, function := range functions {
		params := make([]ApiParam, 0)
		// 构造函数不处理
		if function.Name == clsInfo.NodeName {
			continue
		}

		// 只要 public 的非静态方法
		if !isPublicAndNotStaticMethod(function.Modifiers) {
			continue
		}

		// void 返回值不处理
		if isVoidReturn(function.ReturnType) {
			continue
		}

		// 遍历 function
		for _, param := range function.Parameters {
			params = append(params, ApiParam{
				Name: param.TypeValue,
				Type: paramTypeToGscType(param.TypeType),
			})
		}
		v := ApiAction{
			Name:        function.Name,
			ParamsArray: params,
			Giveback:    function.ReturnType,
			Spec:        buildSpec(function.Annotations),
		}
		if len(v.Spec) > 0 {
			actions = mixActions(actions, v)
			fmt.Printf("[%s] 的值是：%v\n", function.Name, v)
		}
	}

	return ApiDescription{
		Name:    clsInfo.NodeName,
		Actions: actions,
	}, nil
}

func (p *JavaApiDescBuilder) buildApiJson(apiDir string) (map[string]ApiDescription, error) {
	apiDesc := make(map[string]ApiDescription)

	//dir, err := os.Getwd()
	//if err != nil {
	//	println("get current path error")
	//	return nil, err
	//}

	// 遍历当前目录下 ./src/main/java/Api 文件夹下的所有文件
	// apiDir := filepath.Join(dir, path)

	// 遍历当前目录下的文件和子目录
	err := filepath.Walk(apiDir, func(path string, info os.FileInfo, err error) error {
		if apiDir == path {
			return nil
		}
		if err != nil {
			fmt.Printf("遍历目录时出错: %v\n", err)
			return err
		}
		// 不处理目录
		if info.IsDir() {
			return filepath.SkipDir
		}
		clsName, _ := GetFileName(info.Name())
		// 不处理api切片定义文件
		if isProxyApi(clsName) {
			return filepath.SkipDir
		}
		// 不处理非java文件
		if !strings.HasSuffix(info.Name(), ".java") {
			return filepath.SkipDir
		}
		v, err := p.buildApiDescription(path)
		if err != nil {
			fmt.Printf("构建api描述时出错: %v\n", err)
			return err
		}
		if len(v.Actions) > 0 {
			apiDesc[clsName] = v
		}
		return nil
	})
	if err != nil {
		fmt.Printf("遍历目录时出错: %v\n", err)
		return nil, err
	}

	return apiDesc, nil
}

// "src/main/java/Api"

type JavaApiDescBuilder struct {
	root        string
	mavenHelper *ClassApiDescBuilder
}

func (p *JavaApiDescBuilder) BuildApiJson(path string) (map[string]ApiDescription, error) {
	apiDir := filepath.Join(p.root, path)
	return p.buildApiJson(apiDir)
}

func NewApiDescBuilderFromJava(root string) *JavaApiDescBuilder {
	return &JavaApiDescBuilder{root: root, mavenHelper: NewNewApiDescBuilderFromMaven()}
}
