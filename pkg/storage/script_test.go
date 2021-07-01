package storage

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestAddScript(t *testing.T) {
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

	if err := AddScript(script); err != nil {
		t.Fatalf("failed to create new script: %v", err)
	}

	scriptPath := GetScriptDir(script.Name + "/index.sh")
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		t.Fatalf("script file does not exist: %v", err)
	}

	b, err := ioutil.ReadFile(scriptPath)
	if err != nil {
		t.Fatalf("unable to read from script file: %v", err)
	}
	content := string(b)
	if !strings.Contains(content, script.Content) {
		t.Fatalf("script does not contain proper content.\ngot: \n%v\nmust contain:\n%v", content, script.Content)
	}
}

func TestInstallScript(t *testing.T) {
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

	if err := AddScript(script); err != nil {
		t.Fatalf("failed to create new script: %v", err)
	}

	if err := InstallScript(script.Name); err != nil {
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

func TestCreateScriptDir(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Unable to generate a temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if err := Init(dir); err != nil {
		t.Fatalf("failed to initalize dir: %v", err)
	}

	if err := CreateScriptDir("hellotest"); err != nil {
		t.Fatalf("unable to create script dir: %v", err)
	}

	if _, err := os.Stat(GetScriptDir("hellotest")); os.IsNotExist(err) {
		t.Fatalf("failed to find scripts dir: %v", err)
	}

	if _, err := os.Stat(GetScriptDir("hellotest/index.sh")); os.IsNotExist(err) {
		t.Fatalf("failed to find scripts dir: %v", err)
	}

	if _, err := os.Stat(GetScriptDir("hellotest/config.json")); os.IsNotExist(err) {
		t.Fatalf("failed to find scripts dir: %v", err)
	}

}

func TestGetAllScriptDirs(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Unable to generate a temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if err := Init(dir); err != nil {
		t.Fatalf("failed to initalize dir: %v", err)
	}

	scripts := []*Script{
		{Name: "test1", Content: "echo hello"},
		{Name: "test2", Content: "echo hello"},
		{Name: "test2/nested1", Content: "echo hello"},
		{Name: "test2/nested2", Content: "echo hello"},
		{Name: "test2/nested2/nestednested1", Content: "echo hello"},
	}

	for _, script := range scripts {
		AddScript(script)
	}

	scriptDirs, err := GetAllScriptDirs()
	if err != nil {
		t.Fatalf("unable to get all script dirs: %v", err)
	}

	if len(scriptDirs) != len(scripts) {
		t.Fatalf("len(scriptDirs) != len(scripts)\nhave: %v\nwant: %v", len(scriptDirs), len(scripts))
	}

	for _, script := range scripts {
		expectedDir := GetScriptDir(script.Name) + "/index.sh"
		if !contains(scriptDirs, expectedDir) {
			t.Fatalf("Expected to have directory \n%v \nin \n%v", expectedDir, strings.Join(scriptDirs, "\n"))
		}
	}
}

func contains(arr []string, check string) bool {
	for _, el := range arr {
		if check == el {
			return true
		}
	}
	return false
}