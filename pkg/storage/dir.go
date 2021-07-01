package storage

import (
	"fmt"
	"os"
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
