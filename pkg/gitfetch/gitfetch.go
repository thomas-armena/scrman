package gitfetch

import (
	"fmt"
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/thomas-armena/scrman/pkg/storage"
)

type Repository struct {
	Author string
	Name   string
}

func FetchRepo(repo Repository) error {

	scrmanDir := storage.GetRootDir()
	scriptDir := fmt.Sprintf("%v/scripts/%v/%v", scrmanDir, repo.Author, repo.Name)
	if err := os.MkdirAll(scriptDir, 0777); err != nil {
		return fmt.Errorf("unable to fetch repo %v: %v", repo.Name, err)
	}
	url := fmt.Sprintf("https://github.com/%v/%v", repo.Author, repo.Name)

	_, err := git.PlainClone(scriptDir, false, &git.CloneOptions{
		URL:      url,
		Progress: os.Stdout,
	})

	if err != nil {
		return fmt.Errorf("unable to fetch repo %v: %v", repo.Name, err)
	}

	err = allowPermissionsForScripts(scriptDir)
	if err != nil {
		return fmt.Errorf("unable to fetch repo %v: %v", repo.Name, err)
	}

	return nil
}

func allowPermissionsForScripts(root string) error {

	allScriptDirs, err := storage.GetAllScriptDirs()
	if err != nil {
		return err
	}

	for _, path := range allScriptDirs {
		os.Chmod(path, 0777)
	}
	return nil
}
