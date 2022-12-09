package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		msg := fmt.Sprintf("Usage: %v <path_to_directory>", args[0])
		log.Fatal(msg)
	}

	fullPath := args[1]
	dirMap := make(map[string]uint32)

}
