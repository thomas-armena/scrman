package run

import (
	"fmt"
	"os"
	"os/exec"
	"scr/config"
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

	scriptArgs := make([]string, 0)
	scriptArgs = append(scriptArgs, "")
	config, err := config.GetConfig(scriptName)
	if err != nil {
		return err
	}
	for _, arg := range config.Arguments {
		var input string
		fmt.Printf("%v (default: %v): ", arg.Description, arg.Default)
		fmt.Scanln(&input)
		if input == "" {
			input = arg.Default
		}
		scriptArgs = append(scriptArgs, input)
	}

	cmd := &exec.Cmd{
		Path:   scriptDir + "/index.sh",
		Args:   scriptArgs,
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
