package dir

import (
	"fmt"
	"os"
)

func GetScriptDir(scriptName string) (string, error) {
	scrmanDir, err := getScrmanDir()
	if err != nil {
		return "", err
	}
	return scrmanDir + "/scripts/" + scriptName, nil
}

func GetBinDir() (string, error) {
	scrmanDir, err := getScrmanDir()
	if err != nil {
		return "", err
	}
	return scrmanDir + "/bin/", nil
}

func getScrmanDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("unable to get home dir: %v", err)
	}
	return homeDir + "/.scrman", nil
}
