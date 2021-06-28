package install

import (
	"html/template"
	"os"
	"scr/dir"

	"github.com/gobuffalo/packr"
)

type InstallScript struct {
	ScriptName string
}

func Install(args []string) error {
	scriptName := args[0]
	// TODO: Validate scriptName exists

	box := packr.NewBox("../templates/")
	scriptTemplateText, err := box.FindString("install.sh")
	if err != nil {
		return err
	}

	script := &InstallScript{ScriptName: scriptName}

	scriptTemplate, err := template.New("script").Parse(scriptTemplateText)
	if err != nil {
		return err
	}

	binDirectory, err := dir.GetBinDir()
	if err != nil {
		return err
	}
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
