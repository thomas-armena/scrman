package templates

import "github.com/gobuffalo/packr"

var box packr.Box = packr.NewBox("./templates")

func Find(path string) ([]byte, error) {
	return box.Find(path)
}

func FindString(path string) (string, error) {
	return box.FindString(path)
}
