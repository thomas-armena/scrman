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

func TestGetScriptSubDirs(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Unable to generate a temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if err := Init(dir); err != nil {
		t.Fatalf("failed to initalize dir: %v", err)
	}

	scripts := []*Script{
		{Name: "node-a", Content: ""},
		{Name: "node-b", Content: ""},
		{Name: "node-b/node-b-a", Content: ""},
		{Name: "node-b/node-b-b/node-b-b-a", Content: ""},
		{Name: "node-b/node-b-b/node-b-b-b", Content: ""},
		{Name: "node-c", Content: ""},
		{Name: "node-d", Content: ""},
	}

	for _, script := range scripts {
		AddScript(script)
	}

	subdirs, err := GetScriptSubDirs("node-b")
	expected := []string{"node-b/node-b-a", "node-b/node-b-b"}
	if err != nil {
		t.Fatalf("failed to get all script subdirs: %v", err)
	}
	if !compare(subdirs, expected) {
		t.Errorf("incorrect subdirs \ngot: %v\nexpected: %v", subdirs, expected)
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
		expectedDir := GetScriptDir(script.Name)
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

func compare(arr1 []string, arr2 []string) bool {
	if len(arr1) != len(arr2) {
		return false
	}
	for i := range arr1 {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}
