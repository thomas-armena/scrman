package run

import (
	"fmt"
	"os"
	"os/exec"
	"scr/dir"
)

func Run(args []string) error {
	scriptName := args[0]
	err := runScript(scriptName)
	if err != nil {
		return fmt.Errorf("run: Unable to run script %v: %v", scriptName, err)
	}
	return nil
}

func runScript(scriptName string) error {
	scriptDir, err := dir.GetScriptDir(scriptName)
	if err != nil {
		return err
	}

	cmd := &exec.Cmd{
		Path:   scriptDir + "/index.sh",
		Args:   []string{},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
