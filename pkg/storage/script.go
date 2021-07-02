package storage

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/thomas-armena/scrman/pkg/templates"
)

type Script struct {
	Name    string
	Content string
}

func AddScript(script *Script) error {
	if err := CreateScriptDir(script.Name); err != nil {
		return fmt.Errorf("unable to create script dir: %v", err)
	}

	scriptDir := GetScriptDir(script.Name)
	indexFile, err := os.OpenFile(scriptDir+"/index.sh", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("create: %v", err)
	}
	defer indexFile.Close()
	if _, err = indexFile.WriteString(script.Content); err != nil {
		return err
	}
	fmt.Println("Script created in: " + scriptDir)

	return nil
}

func InstallScript(scriptName string) error {

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

func CreateScriptDir(scriptName string) error {
	if scriptName == "" {
		return fmt.Errorf("script name cannot be empty")
	}

	scriptDir := GetScriptDir(scriptName)
	if err := os.MkdirAll(scriptDir, 0777); err != nil {
		return fmt.Errorf("unable to create new script directory: %v", err)
	}

	index, err := templates.Find("script/index.sh")
	if err != nil {
		return fmt.Errorf("unable to find index.sh template: %v", err)
	}

	config, err := templates.Find("script/config.json")
	if err != nil {
		return fmt.Errorf("unable to find config.json template: %v", err)
	}

	scriptPath := scriptDir + "/index.sh"
	if err := os.WriteFile(scriptPath, index, 0777); err != nil {
		return fmt.Errorf("unable to write index.sh: %v", err)
	}

	configPath := scriptDir + "/config.json"
	if err := os.WriteFile(configPath, config, 0777); err != nil {
		return fmt.Errorf("unable to write config.json: %v", err)
	}

	return nil
}

func GetAllScriptDirs() ([]string, error) {
	scriptDirs := make([]string, 0)

	err := filepath.Walk(scrmanRoot,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if getLeafOfPath(path) == "index.sh" {
				scriptDirs = append(scriptDirs, path)
			}
			return nil
		})
	if err != nil {
		return scriptDirs, err
	}
	return scriptDirs, nil
}

func getLeafOfPath(path string) string {
	separated := strings.Split(path, "/")
	return separated[len(separated)-1]
}
