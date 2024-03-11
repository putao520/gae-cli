package api_client

import (
	"errors"
	"fmt"
	parser "gae-cli/gsc/wreulicke/classfile-parser"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// 非 static 的 public 方法
func filePublicAndNotStatic(method *parser.Method) bool {
	// public = 1
	// static = 8
	return method.AccessFlags == 1
}

func classPrueName(name string) string {
	if strings.HasPrefix(name, "L") {
		return name[1 : len(name)-1]
	}
	switch name {
	case "I":
		return "int"
	case "F":
		return "float"
	case "D":
		return "double"
	case "C":
		return "char"
	case "J":
		return "long"
	case "Z":
		return "boolean"
	case "V":
		return "void"
	case "S":
		return "short"
	case "B":
		return "byte"
	default:
		if strings.HasSuffix(name, "[") {
			panic("api return type not support array type: " + name)
		}
		if strings.HasSuffix(name, "(") {
			panic("api return type not support function type: " + name)
		}
		panic("unknown type: " + name)
	}
}

func buildApiParamForClass(method *parser.Method, desc string, constantPool *parser.ConstantPool) ([]ApiParam, string, error) {
	result := make([]ApiParam, 0)
	// 获取 () 之间的内容
	// 在 desc 中查找 ( ,当大于1个报错
	if strings.Count(desc, "(") > 1 {
		// panic("method not support function type param or return: " + desc)
		return nil, "", errors.New("method not support function type param or return: " + desc)
	}
	methodParamArray := method.MethodParameters()

	c1 := strings.Index(desc, "(")
	c2 := strings.Index(desc, ")")
	if c1 == -1 || c2 == -1 {
		return nil, "", errors.New("method desc is invalid: " + desc)
	}

	i := 0
	parameters := desc[c1+1 : c2]
	for len(parameters) > 0 {
		param := parameters
		// L 开头代表对象,查找;作为分割
		if strings.HasPrefix(parameters, "L") {
			// 逐字符遍历 parameters
			for u, c := range parameters {
				if c == ';' {
					param = parameters[1:u]
					parameters = parameters[u+1:]
					break
				}
			}
		} else {
			param = classPrueName(parameters[:1])
			parameters = parameters[1:]
		}

		paramName := "Param_" + strconv.Itoa(i)
		if methodParamArray != nil {
			paramNameUtf8, err := constantPool.GetConstantUtf8(methodParamArray.Parameters[i].NameIndex)
			if err == nil {
				paramName = paramNameUtf8.String()
			}
		}
		// 分割参数
		paramArray := strings.Split(param, "/")
		// 获取最后一个元素
		paramType := paramArray[len(paramArray)-1]

		result = append(result, ApiParam{
			Name: paramName,
			Type: paramTypeToGscType(paramType),
		})

		i++
	}

	// 取 ) 后的内容
	li := strings.LastIndex(desc, ")")
	giveback := desc[li+1:]
	// 去掉 L 和 ;
	giveback = classPrueName(giveback)
	// 分割参数
	gArr := strings.Split(giveback, "/")
	if gArr[0] == "java" && gArr[1] == "lang" {
		gArr = gArr[2:]
	}
	giveback = strings.Join(gArr, ".")
	return result, giveback, nil
}

func mapApiSpec(name string) ApiSpec {
	switch name {
	case "PublicApi":
		return PublicApi
	case "SessionApi":
		return SessionApi
	case "OauthApi":
		return OauthApi
	case "PrivateApi":
		return PrivateApi
	case "CloseApi":
		return CloseApi
	default:
		return ""
	}
}

// 过滤注解,按GSC的注解规则过滤
func filterApiAnnotation(method *parser.Method, c *parser.ConstantPool) []ApiSpec {
	specs := make([]ApiSpec, 0)
	visAnnotations := method.RuntimeVisibleAnnotations()
	if visAnnotations != nil {
		annotations := visAnnotations.Annotations
		for _, annotation := range annotations {
			annotationFullName, err := annotation.Type(c)
			if err != nil {
				println("get annotation type error")
				continue
			}
			// Lcommon/java/InterfaceModel/Type/InterfaceType;
			if annotationFullName != "Lcommon/java/InterfaceModel/Type/InterfaceType;" {
				continue
			}
			eleArray := annotation.ElementValuePairs
			for _, ele := range eleArray {
				if constValue, ok := ele.ElementValue.(*parser.ElementValueEnumConstValue); ok {
					//typeName, err := c.GetConstantUtf8(constValue.TypeNameIndex)
					//if err != nil {
					//	println("get type name error")
					//	continue
					//}
					//println(typeName.String())
					constName, err := c.GetConstantUtf8(constValue.ConstNameIndex)
					if err != nil {
						println("get const name error")
						continue
					}
					// println(constName.String())
					apiTypeName := constName.String()
					if apiTypeName == "CloseApi" || apiTypeName == "PrivateApi" {
						return nil
					}
					specs = append(specs, mapApiSpec(apiTypeName))
				}
			}
		}
	}
	if len(specs) == 0 {
		specs = append(specs, PublicApi)
	}
	return specs
}

func (t *ClassApiDescBuilder) BuildApiJson(classFilePath string) (ApiDescription, error) {
	apiDir := filepath.Join(t.root, classFilePath)
	f, err := os.Open(apiDir)
	if err != nil {
		println("open class file[ " + apiDir + " ] error")
		return ApiDescription{}, err
	}
	p := parser.New(f)
	classFileParser, err := p.Parse()
	if err != nil {
		println("parse class file[ " + apiDir + " ] error")
		return ApiDescription{}, err
	}
	defer f.Close()

	actions := make([]ApiAction, 0)

	name, err := classFileParser.ThisClassName()
	if err != nil {
		println("get class name error")
		return ApiDescription{}, err
	}

	// 处理父类
	fatherName, err := classFileParser.SuperClassName()
	if err != nil {
		println("get father class name error")
		return ApiDescription{}, err
	}
	if fatherName != "java/lang/Object" && fatherName != "" {
		// 解析父类
		apiDesc, err := t.BuildApiJson(fatherName)
		if err != nil {
			println("build father class error")
			return ApiDescription{}, err
		}
		actions = append(actions, apiDesc.Actions...)
	}

	c := classFileParser.ConstantPool
	for _, method := range classFileParser.Methods {
		if !filePublicAndNotStatic(method) {
			continue
		}
		methodName, err := method.Name(c)

		println("methodName: " + methodName)

		// 过滤注解
		spec := filterApiAnnotation(method, c)
		if spec == nil {
			// 过滤 CloseApi 和 PrivateApi 注解的方法
			continue
		}

		// 过滤构造方法和静态代码块
		if methodName == "<init>" || methodName == "<clinit>" {
			continue
		}
		if err != nil {
			println("get method name error")
			return ApiDescription{}, err
		}

		// 处理参数
		methodDesc, err := method.Descriptor(c)
		if err != nil {
			println("get method desc error")
			return ApiDescription{}, err
		}

		params, giveback, err := buildApiParamForClass(method, methodDesc, c)
		if err != nil {
			println("build api param error")
			continue
		}
		v := ApiAction{
			Name:        methodName,
			ParamsArray: params,
			Giveback:    giveback,
			Spec:        spec,
		}
		if len(v.Spec) > 0 {
			actions = append(actions, v)
			fmt.Printf("[%s] 的值是：%v\n", methodName, v)
		}
	}

	return ApiDescription{
		Name:    name,
		Actions: actions,
	}, nil
}

type ClassApiDescBuilder struct {
	root string
}

func NewApiDescBuilderFromClass(root string) *ClassApiDescBuilder {
	return &ClassApiDescBuilder{
		root: root,
	}
}

func NewNewApiDescBuilderFromMaven() *ClassApiDescBuilder {
	tmpDir, err := LoadPomDependencyPackage()
	if err != nil {
		println("load pom dependencies package error")
	}
	//defer func() {
	//	// 删除临时文件夹
	//	if err == nil {
	//		// err := os.RemoveAll(tmpDir)
	//		if err != nil {
	//			return
	//		}
	//	}
	//}()
	return &ClassApiDescBuilder{
		root: tmpDir,
	}
}
