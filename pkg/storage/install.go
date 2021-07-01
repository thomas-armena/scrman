package storage

import (
	"html/template"
	"os"
	"strings"

	"github.com/thomas-armena/scrman/pkg/templates"
)

func Install(scriptName string) error {

	scriptTemplateText, err := templates.FindString("install.sh")
	if err != nil {
		return err
	}

	scriptTemplate, err := template.New("script").Parse(scriptTemplateText)
	if err != nil {
		return err
	}

	binDirectory := GetBinDir()
	binaryName := getLeafOfPath(scriptName)
	file, err := os.Create(binDirectory + "/" + binaryName)

	if err != nil {
		return err
	}
	err = file.Chmod(0777)
	if err != nil {
		return err
	}
	scriptTemplate.Execute(file, scriptName)

	return nil
}

func getLeafOfPath(path string) string {
	separated := strings.Split(path, "/")
	return separated[len(separated)-1]
}
