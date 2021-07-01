package storage

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("Unable to generate a temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	if err := Init(dir); err != nil {
		t.Fatalf("failed to initalize dir: %v", err)
	}

	testCases := []struct {
		name     string
		expected string
		got      string
	}{
		{name: "root", expected: dir, got: GetRootDir()},
		{name: "scripts", expected: dir + "/scripts/", got: GetScriptDir("")},
		{name: "scripts/hello", expected: dir + "/scripts/hello", got: GetScriptDir("hello")},
		{name: "bin", expected: dir + "/bin/", got: GetBinDir()},
	}

	for _, tc := range testCases {
		if tc.expected != tc.got {
			t.Fatalf("%v = %q; want %q", tc.name, tc.got, tc.expected)
		}
	}

	if _, err := os.Stat(GetRootDir()); os.IsNotExist(err) {
		t.Fatalf("failed to find root dir: %v", err)
	}

	if _, err := os.Stat(GetScriptDir("")); os.IsNotExist(err) {
		t.Fatalf("failed to find scripts dir: %v", err)
	}

	if _, err := os.Stat(GetBinDir()); os.IsNotExist(err) {
		t.Fatalf("failed to find bin dir: %v", err)
	}
}
