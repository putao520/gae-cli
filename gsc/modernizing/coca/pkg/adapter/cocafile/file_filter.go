package cocafile

import (
	"path/filepath"
	"strings"
)

var isJavaTestFile = func(path string) bool {
	return strings.HasSuffix(path, "Test.java") || strings.HasSuffix(path, "Tests.java")
}

var isJavaTestPackage = func(path string) bool {
	path = filepath.ToSlash(path)
	return strings.Contains(path, "src/test/java/")
}

var JavaTestFileFilter = func(path string) bool {
	return isJavaTestFile(path) || isJavaTestPackage(path)
}

var JavaCodeFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".java") && !JavaTestFileFilter(path)
}

var JavaFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".java")
}

var TypeScriptFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".ts")
}

var PythonFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".py")
}

var GoFileFilter = func(path string) bool {
	return strings.HasSuffix(path, ".go")
}

var PomXmlFilter = func(path string) bool {
	return strings.HasSuffix(path, "pom.xml")
}

var BuildGradleFilter = func(path string) bool {
	return strings.HasSuffix(path, "build.gradle")
}
