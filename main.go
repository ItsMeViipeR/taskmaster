package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println(errors.New("Please specify a command"))
		return
	}

	switch args[1] {
	case "add":
		fmt.Println("Add")
	case "remove":
		fmt.Println("Remove")
	case "done":
		fmt.Println("Done")
	case "idof":
		fmt.Println("Idof")
	}
}
