package install

import (
	"fmt"
	"html/template"
	"os"

	"github.com/gobuffalo/packr"
	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/dir"
)

type InstallScript struct {
	ScriptName string
}

func NewCmdInstall() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install an existing bash script",
		RunE: func(cmd *cobra.Command, args []string) error {
			return install(args)
		},
	}

	return cmd
}

func install(args []string) error {
	scriptName := args[0]
	err := InstallByScriptName(scriptName)
	if err != nil {
		return fmt.Errorf("unable to install script: %v", err)
	}
	// TODO: Validate scriptName exists
	return nil
}

// TODO: Move this function to another package. It is also used by create cmd.
func InstallByScriptName(scriptName string) error {
	box := packr.NewBox("../../../templates/")
	scriptTemplateText, err := box.FindString("install.sh")
	if err != nil {
		return err
	}

	script := &InstallScript{ScriptName: scriptName}

	scriptTemplate, err := template.New("script").Parse(scriptTemplateText)
	if err != nil {
		return err
	}

	binDirectory := dir.GetBinDir()

	file, err := os.Create(binDirectory + "/" + scriptName)
	if err != nil {
		return err
	}
	err = file.Chmod(0777)
	if err != nil {
		return err
	}
	scriptTemplate.Execute(file, script)

	return nil
}
