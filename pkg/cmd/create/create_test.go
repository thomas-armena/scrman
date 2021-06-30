package create

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/thomas-armena/scrman/pkg/dir"
)

func TestCreate(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to generate temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := dir.Init(tempDir); err != nil {
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

	if err := create(script, nil); err != nil {
		t.Fatalf("failed to create new script: %v", err)
	}

	scriptPath := dir.GetScriptDir(script.Name + "/index.sh")
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
