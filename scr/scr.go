package main

import (
	"log"
	"os"
	"scr/create"
	"scr/dir"
	"scr/install"
	"scr/run"
)

func main() {

	err := dir.InitDirectories()
	if err != nil {
		log.Fatal(err)
	}

	err = dir.InitProject("helloworld")
	if err != nil {
		log.Fatal(err)
	}

	args := os.Args[1:]

	cmd := args[0]
	cmdArgs := args[1:]

	if cmd == "create" {
		err = create.Create(cmdArgs)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cmd == "run" {
		err = run.Run(cmdArgs)
		if err != nil {
			log.Fatal(err)
		}
	}

	if cmd == "install" {
		err = install.Install(cmdArgs)
		if err != nil {
			log.Fatal(err)
		}
	}

}
