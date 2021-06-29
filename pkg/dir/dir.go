package dir

import (
	"fmt"
	"os"

	"github.com/thomas-armena/scrman/pkg/templates"
)

var scrmanRoot = ""

func DefaultRoot() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get home dir: %v", err)
	}

	root := homeDir + "/.scrman"

	return root, nil
}

func GetRootDir() string {
	return scrmanRoot
}

func GetScriptDir(scriptName string) string {
	return scrmanRoot + "/scripts/" + scriptName
}

func GetBinDir() string {
	return scrmanRoot + "/bin/"
}

func InitDefault() error {
	root, err := DefaultRoot()
	if err != nil {
		return err
	}

	return Init(root)
}

func Init(appRoot string) error {
	scrmanRoot = appRoot

	if err := os.MkdirAll(GetRootDir(), 0777); err != nil {
		return fmt.Errorf("unable to make scrman root directory: %v", err)
	}

	if err := os.MkdirAll(GetScriptDir(""), 0777); err != nil {
		return fmt.Errorf("unable to make scrman scripts directory: %v", err)
	}

	if err := os.MkdirAll(GetBinDir(), 0777); err != nil {
		return fmt.Errorf("unable to make scrman bin directory: %v", err)
	}

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
