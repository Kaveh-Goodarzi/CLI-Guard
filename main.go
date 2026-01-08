package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  guard init <path>")
		fmt.Println("  guard verify <path>")
		os.Exit(1)
	}

	command := os.Args[1]
	path := os.Args[2]

	switch command {
		case "init":
			err := Init(path)
			if err != nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
			fmt.Println("Initialized .guard.json")
		case "verify":
			err := Verify(path)
			if err!= nil {
				fmt.Println("Error: ", err)
				os.Exit(1)
			}
		default:
			fmt.Println("Unknown command", command)
			os.Exit(1)
	}
}