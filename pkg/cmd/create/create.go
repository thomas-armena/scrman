package create

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/chroma/quick"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/storage"
)

const (
	maxCommands        = 25
	maxCommandsPerPage = 10
)

func NewCmdCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new bash script from bash history",
		RunE: func(cmd *cobra.Command, args []string) error {
			script, err := getScriptInput()
			if err != nil {
				return fmt.Errorf("unable to create script dir: %v", err)
			}

			return storage.AddScript(script)
		},
	}

	return cmd
}

func getScriptInput() (*storage.Script, error) {
	history, err := getZshHistory()
	if err != nil {
		return nil, fmt.Errorf("create: %v", err)
	}

	numCommandsToShow := min(maxCommands, len(history))
	pageLength := min(maxCommandsPerPage, numCommandsToShow)
	lastIndex := len(history) - 1

	commands := history[lastIndex-numCommandsToShow : lastIndex]

	startSelectTemplates := &promptui.SelectTemplates{
		Label:    "Select starting command",
		Selected: " ",
	}

	startPrompt := promptui.Select{
		Label:     "Select starting command",
		Items:     commands,
		Size:      pageLength,
		Templates: startSelectTemplates,
	}

	startingIndex, _, err := startPrompt.RunCursorAt(len(commands)-1, len(commands)-pageLength)
	if err != nil {
		return nil, fmt.Errorf("unable to create script dir: %v", err)
	}

	commands = commands[startingIndex:]

	endSelectTemplates := &promptui.SelectTemplates{
		Label:    "Select ending command",
		Selected: "Script:",
	}
	endPrompt := promptui.Select{
		Label:     "Select ending command",
		Items:     commands,
		Size:      pageLength,
		Templates: endSelectTemplates,
	}

	endingIndex, _, err := endPrompt.RunCursorAt(len(commands)-1, len(commands)-1-pageLength)
	if err != nil {
		return nil, fmt.Errorf("unable to create script dir: %v", err)
	}
	commands = commands[:endingIndex+1]

	script := &storage.Script{
		Content: strings.Join(commands, "\n"),
	}

	fmt.Print("═══════════════════════════════\n\n")
	if err := quick.Highlight(os.Stdout, script.Content, "bash", "terminal256", ""); err != nil {
		return nil, fmt.Errorf("unable to highlight script content: %v", err)
	}
	fmt.Print("\n\n═══════════════════════════════\n")

	fmt.Printf("Enter name of script: ")
	fmt.Scanln(&script.Name)

	return script, nil
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
		history = append(history, strings.Trim(x[1], "\n "))
	}
	return history, nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
