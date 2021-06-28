package main

import (
	"log"
	"os"
)

func main() {
	print("Hello World\n")
	err := initDirectories()
	if err != nil {
		log.Fatal(err)
	}
}

func initDirectories() error {

	currDir, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(currDir)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	os.Chdir(homeDir)

	err = os.MkdirAll(".scrman", 0755)
	if err != nil {
		return err
	}

	os.Chdir(".scrman")

	err = os.MkdirAll("bin", 0755)
	if err != nil {
		return err
	}

	err = os.MkdirAll("scripts", 0755)
	if err != nil {
		return err
	}

	return nil
}
