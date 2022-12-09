package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		msg := fmt.Sprintf("Usage: %v <path_to_directory>", args[0])
		log.Fatal(msg)
	}

	fullPath := args[1]
	dirMap := make(map[string]uint32)

	filepath.WalkDir(fullPath, walk)

}

func walk(path string, dir fs.DirEntry, err error) error {
	if err != nil {
		return err
	}

}
