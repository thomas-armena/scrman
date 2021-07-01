package storage

import (
	"fmt"
	"os"
)

type Script struct {
	Name    string
	Content string
}

func Create(script *Script) error {
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
