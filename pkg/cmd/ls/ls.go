package ls

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/storage"
)

func NewCmdLs() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ls",
		Short: "List all script directories",
		RunE: func(cmd *cobra.Command, args []string) error {
			err := listScriptDirs()
			if err != nil {
				return err
			}
			return nil
		},
	}
	return cmd
}

func listScriptDirs() error {
	allScriptDirs, err := storage.GetAllScriptDirs()
	if err != nil {
		return fmt.Errorf("unable to list script directories: %v", err)
	}
	for _, scriptDir := range allScriptDirs {
		fmt.Println(scriptDir)
	}
	return nil
}
