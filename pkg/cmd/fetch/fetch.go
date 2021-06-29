package fetch

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/gitfetch"
)

func NewCmdFetch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch",
		Short: "Fetch a script repository using git",
		RunE: func(cmd *cobra.Command, args []string) error {
			return fetch(args)
		},
	}

	return cmd
}

func fetch(args []string) error {
	repoPath := args[0]
	splitPath := strings.Split(repoPath, "/")
	author := splitPath[0]
	name := splitPath[1]
	repo := gitfetch.Repository{Author: author, Name: name}

	err := gitfetch.FetchRepo(repo)
	if err != nil {
		return err
	}
	return nil
}
