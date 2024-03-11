package api_client

import "strings"

type ApiSpec string

const (
	PublicApi  = "PublicApi"
	SessionApi = "SessionApi"
	OauthApi   = "OauthApi"
	PrivateApi = "PrivateApi"
	CloseApi   = "CloseApi"
)

type ApiParam struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type ApiAction struct {
	Name        string     `json:"name"`
	ParamsArray []ApiParam `json:"paramsArray"`
	Giveback    string     `json:"giveback"`
	Spec        []ApiSpec  `json:"spec"`
}

type ApiDescription struct {
	Name    string      `json:"name"`
	Actions []ApiAction `json:"actions"`
}

func isProxyApi(name string) bool {
	// name 是 After 结尾 或者 name 是 Before 结尾
	if strings.HasSuffix(name, "After") || strings.HasSuffix(name, "Before") {
		return true
	}
	return false
}

func paramTypeToGscType(paramType string) string {
	switch paramType {
	case "int", "Integer":
		return "i"
	case "float", "Float":
		return "f"
	case "double", "Double":
		return "d"
	case "char", "Character":
		return "c"
	case "String":
		return "s"
	case "long", "Long":
		return "l"
	case "boolean", "Boolean":
		return "b"
	case "short", "Short":
		return "short"
	case "JSONObject":
		return "j"
	case "JSONArray":
		return "ja"
	default:
		return "o"
	}
}

func isPublicAndNotStaticMethod(modifiers []string) bool {
	for _, modifier := range modifiers {
		if modifier == "static" || modifier == "private" || modifier == "protected" || modifier == "default" || modifier == "final" || modifier == "abstract" || modifier == "synchronized" || modifier == "native" || modifier == "strictfp" {
			return false
		}
	}
	return true
}

func isVoidReturn(returnType string) bool {
	return returnType == "void"
}

func getSpecStr(annotationStr string) ApiSpec {
	sArr := strings.Split(annotationStr, ".")
	annotationVal := sArr[len(sArr)-1]
	switch annotationVal {
	case "SessionApi":
		return SessionApi
	case "OauthApi":
		return OauthApi
	case "PrivateApi":
		return PrivateApi
	case "CloseApi":
		return CloseApi
	default:
		return PublicApi
	}
}

func Package2Path(javaPath string) string {
	// 将 groupId 转换为路径, . 替换为路径分隔符
	return strings.Join(strings.Split(javaPath, "."), "/")
}
