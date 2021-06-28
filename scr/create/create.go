package create

import "fmt"

func Create(args []string) error {
	fmt.Println("Running init")
	fmt.Println(args)
	return nil
}
