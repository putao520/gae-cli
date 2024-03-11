package api_client

import (
	"github.com/beevik/etree"
	"os"
)

type MavenDependency struct {
	GroupId    string
	ArtifactId string
	Version    string
	Packaging  string // war Or jar
}

func LoadPomDependencies() ([]MavenDependency, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		println("get current path error")
		return nil, err
	}
	pomPath := currentPath + "/pom.xml"
	pomContent, err := os.ReadFile(pomPath)
	if err != nil {
		println("read pom.xml error")
		return nil, err
	}
	doc := etree.NewDocument()
	if err := doc.ReadFromBytes(pomContent); err != nil {
		panic(err)
	}
	dependencies := doc.FindElements("//dependencies/dependency")
	var result []MavenDependency
	for _, dependency := range dependencies {
		groupId := dependency.FindElement("groupId")
		artifactId := dependency.FindElement("artifactId")
		version := dependency.FindElement("version")
		packaging := dependency.FindElement("packaging")
		if packaging == nil {
			packaging = etree.NewElement("packaging")
			packaging.SetText("jar")
		}
		result = append(result, MavenDependency{
			GroupId:    groupId.Text(),
			ArtifactId: artifactId.Text(),
			Version:    version.Text(),
			Packaging:  packaging.Text(),
		})
	}
	return result, nil
}
