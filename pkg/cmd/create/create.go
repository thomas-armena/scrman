package create

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/cmd/install"
	"github.com/thomas-armena/scrman/pkg/dir"
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new bash script from bash history",
		RunE: func(cmd *cobra.Command, args []string) error {
			return create(args)
		},
	}

	return cmd
}

func create(args []string) error {

	history, err := getZshHistory()
	if err != nil {
		return fmt.Errorf("create: %v", err)
	}
	recentHistory := history[len(history)-11 : len(history)-1]
	for i, cmd := range recentHistory {
		fmt.Printf("%v: %v", i+1, cmd)
	}

	fmt.Printf("Enter number of command to start at: ")
	var startingIndex int
	fmt.Scan(&startingIndex)

	fmt.Printf("Enter name of script: ")
	var scriptName string
	fmt.Scanln(&scriptName)

	scriptText := strings.Join(recentHistory[startingIndex-1:], "")
	fmt.Println(scriptText)

	err = dir.InitProject(scriptName)
	if err != nil {
		return fmt.Errorf("create: init project: %v", err)
	}

	scriptDir, err := dir.GetScriptDir(scriptName)
	if err != nil {
		return fmt.Errorf("create: %v", err)
	}
	indexFile, err := os.OpenFile(scriptDir+"/index.sh", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		return fmt.Errorf("create: %v", err)
	}
	defer indexFile.Close()
	if _, err = indexFile.WriteString(scriptText); err != nil {
		return err
	}

	err = install.InstallByScriptName(scriptName)
	if err != nil {
		return fmt.Errorf("create: %v", err)
	}

	return nil
}

// Assumes history is stored in ~/.zsh_history
func getZshHistory() ([]string, error) {
	history := make([]string, 0)
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return history, fmt.Errorf("unable to get history: %v", err)
	}
	file, err := os.ReadFile(homeDir + "/.zsh_history")
	if err != nil {
		return history, fmt.Errorf("unable to get history: %v", err)
	}
	historyString := string(file)
	lines := strings.Split(historyString, ":")
	for _, line := range lines {
		x := strings.Split(line, ";")
		if len(x) < 2 {
			continue
		}
		history = append(history, x[1])
	}
	return history, nil
}
