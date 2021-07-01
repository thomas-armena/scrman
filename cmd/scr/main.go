package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/thomas-armena/scrman/pkg/cmd/create"
	"github.com/thomas-armena/scrman/pkg/cmd/fetch"
	"github.com/thomas-armena/scrman/pkg/cmd/install"
	"github.com/thomas-armena/scrman/pkg/cmd/run"
	"github.com/thomas-armena/scrman/pkg/storage"
)

func main() {
	if err := storage.InitDefault(); err != nil {
		log.Fatalf("Unable to initalize scr directories: %v", err)
	}

	rootCmd := &cobra.Command{
		Use:   "scr",
		Short: "A Unix commands line interface for bash script managing and sharing",
	}

	createCmd := create.NewCmdCreate()
	runCmd := run.NewCmdRun()
	installCmd := install.NewCmdInstall()
	fetchCmd := fetch.NewCmdFetch()

	rootCmd.AddCommand(
		createCmd,
		runCmd,
		installCmd,
		fetchCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Unable to execute scr: %v", err)
	}

}
