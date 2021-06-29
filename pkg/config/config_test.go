package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/thomas-armena/scrman/pkg/dir"
)

func TestGetConfig(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("unable to generate temp dir: %v", err)
	}
	defer os.RemoveAll(tempDir)

	if err := dir.Init(tempDir); err != nil {
		t.Fatalf("unable to initalize dir: %v", err)
	}

	if _, err := GetConfig("FakeScript"); err == nil {
		t.Fatalf(`GetConfig("FakeScript").err  = nil; want error`)
	}

	if err := dir.CreateScriptDir("helloTest"); err != nil {
		t.Fatalf("unable to create script dir: %v", err)
	}

	config, err := GetConfig("helloTest")
	if err != nil {
		t.Fatalf("failed to get existing config.json: %v", err)
	}

	if config.Location != "./" {
		t.Fatalf("config.Location = %q; want %q", config.Location, "./")
	}

	if len(config.Arguments) > 0 {
		t.Fatalf("len(config.Arguments) > 0; want 0")
	}

	testConfig := `{
	"location": "./some/path/test",
	"arguments": [
		{
			"description": "123", 
			"default": "1"
		},
		{
			"description": "abc", 
			"default": "2"
		}
	]
}
`
	configPath := dir.GetScriptDir("helloTest/config.json")
	if err := ioutil.WriteFile(configPath, []byte(testConfig), 0777); err != nil {
		t.Fatalf("unable to write test config to config.json: %v", err)
	}

	config, err = GetConfig("helloTest")
	if err != nil {
		t.Fatalf("failed to get config.json: %v", err)
	}

	if config.Location != "./some/path/test" {
		t.Fatalf("config.Location = %q; want %q", config.Location, "./some/path/test")
	}

	if len(config.Arguments) != 2 {
		t.Fatalf("len(config.Arguments) != 2; want 2")
	}

	testCases := []struct {
		Description string
		Default     string
	}{
		{Description: "123", Default: "1"},
		{Description: "abc", Default: "2"},
	}

	for i, tc := range testCases {
		if tc.Description != config.Arguments[i].Description {
			t.Fatalf("config.Arguments[%d].Description = %v; want %v", i, config.Arguments[i].Description, tc.Description)
		}

		if tc.Default != config.Arguments[i].Default {
			t.Fatalf("config.Arguments[%d].Default = %v; want %v", i, config.Arguments[i].Default, tc.Default)
		}
	}
}
