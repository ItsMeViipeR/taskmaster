package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/ItsMeViipeR/taskmaster/storage"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println(errors.New("Please specify a command"))
		return
	}

	switch args[1] {
	case "add":
		json := storage.NewJSON("tasks.json")

		if load_err := json.Load(); load_err != nil {
			fmt.Println(load_err)
			return
		}
	case "remove":
		fmt.Println("Remove")
	case "done":
		fmt.Println("Done")
	case "idof":
		fmt.Println("Idof")
	}
}
