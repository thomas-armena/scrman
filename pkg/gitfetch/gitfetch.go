package gitfetch

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/go-git/go-git/v5"
	"github.com/thomas-armena/scrman/pkg/dir"
)

type Repository struct {
	Author string
	Name   string
}

func FetchRepo(repo Repository) error {

	scrmanDir := dir.GetRootDir()
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
	err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			os.Chmod(path, 0777)
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	return nil
}
