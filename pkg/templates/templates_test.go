package templates

import (
	"io/fs"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestFind(t *testing.T) {
	err := filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		expected, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}

		got, err := Find(rel)
		if err != nil {
			return err
		}

		if string(expected) != string(got) {
			t.Fatalf("got = %v; want %v", string(got), string(expected))
		}

		return nil
	})

	if err != nil {
		t.Fatalf("failed to find file: %v", err)
	}
}

func TestFindString(t *testing.T) {
	err := filepath.Walk("./templates", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		expected, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}

		rel, err := filepath.Rel("templates", path)
		if err != nil {
			return err
		}

		got, err := FindString(rel)
		if err != nil {
			return err
		}

		if string(expected) != got {
			t.Fatalf("got = %v; want %v", got, string(expected))
		}

		return nil
	})

	if err != nil {
		t.Fatalf("failed to find file: %v", err)
	}
}
