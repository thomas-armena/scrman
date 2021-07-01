package install

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/storage"
)

func NewCmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install an existing bash script",
		RunE: func(cmd *cobra.Command, args []string) error {
			scriptName := args[0]
			err := storage.Install(scriptName)
			if err != nil {
				return fmt.Errorf("unable to install script: %v", err)
			}
			return nil
		},
	}
	return cmd
}
