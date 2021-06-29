package dir

import (
	"fmt"
	"os"

	"github.com/gobuffalo/packr"
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

func InitDirectories() error {

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	os.Chdir(homeDir)

	err = os.MkdirAll(".scrman", 0777)
	if err != nil {
		return err
	}

	os.Chdir(".scrman")

	err = os.MkdirAll("bin", 0777)
	if err != nil {
		return err
	}

	err = os.MkdirAll("scripts", 0777)
	if err != nil {
		return err
	}

	// TODO: Remove this
	err = initHelloWorld()
	if err != nil {
		return err
	}
	return nil
}

func InitProject(projectName string) error {

	// TODO: validate project name

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	projectDir := homeDir + "/.scrman/scripts/" + projectName

	err = os.MkdirAll(projectDir, 0777)
	if err != nil {
		return err
	}

	os.Chdir(projectDir)

	box := packr.NewBox("../../../templates/script")
	index, err := box.Find("index.sh")
	if err != nil {
		return fmt.Errorf("unable to get index.sh: %v", err)
	}

	config, err := box.Find("config.json")
	if err != nil {
		return fmt.Errorf("unable to get config.json: %v", err)
	}

	err = os.WriteFile("index.sh", index, 0777)
	if err != nil {
		return fmt.Errorf("unable to write index.sh: %v", err)
	}

	err = os.WriteFile("config.json", config, 0777)
	if err != nil {
		return fmt.Errorf("unable to write config.json: %v", err)
	}

	return nil
}

// TODO: Remove this
func initHelloWorld() error {
	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	projectDir := homeDir + "/.scrman/scripts/helloworld"

	err = os.MkdirAll(projectDir, 0777)
	if err != nil {
		return err
	}

	os.Chdir(projectDir)

	box := packr.NewBox("../../templates/helloworld")
	index, err := box.Find("index.sh")
	if err != nil {
		return fmt.Errorf("unable to get index.sh: %v", err)
	}

	config, err := box.Find("config.json")
	if err != nil {
		return fmt.Errorf("unable to get config.json: %v", err)
	}

	err = os.WriteFile("index.sh", index, 0777)
	if err != nil {
		return fmt.Errorf("unable to write index.sh: %v", err)
	}

	err = os.WriteFile("config.json", config, 0777)
	if err != nil {
		return fmt.Errorf("unable to write config.json: %v", err)
	}

	return nil
}
