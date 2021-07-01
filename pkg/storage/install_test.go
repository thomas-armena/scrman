package storage

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestInstall(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to generate temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := Init(tempDir); err != nil {
		t.Fatalf("unable to initalize dir: %v", err)
	}

	script := &Script{
		Name: "testscript",
		Content: `
echo "Hello"
echo "World"
echo "asbklnsklnlkanklsnvsaivknlwq"
		`,
	}

	if err := Create(script); err != nil {
		t.Fatalf("failed to create new script: %v", err)
	}

	if err := Install(script.Name); err != nil {
		t.Fatalf("failed to install new script: %v", err)
	}

	b, err := ioutil.ReadFile(GetBinDir() + script.Name)
	if err != nil {
		t.Fatalf("unable to read from script file: %v", err)
	}
	content := string(b)

	requiredContent := "#!/usr/bin/env bash\nscr run " + script.Name

	if !(content == requiredContent) {
		t.Fatalf("install script does not contain proper content.\ngot: \n%v\nwant:\n%v", content, requiredContent)
	}

}
