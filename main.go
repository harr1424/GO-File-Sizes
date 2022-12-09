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

	// dirMap should be passed here
	filepath.WalkDir(fullPath, walk)

}

// how to create paramter accepting &map[string]uint32 ?
func walk(path string, entry fs.DirEntry, err error, directoryMap map[string]uint32) error {
	if err != nil {
		return err
	}

	if !entry.IsDir() {
		// add entry and it's size to dirMap

		file, err := os.Stat(path)
		if err != nil {
			log.Println("Error determining file size:", err)
		}

		size := file.Size()
		directoryMap[path] = uint32(size)

	} else if entry.IsDir() {
		// recurse here

	} else {
		log.Println("Unhandled branch:")
		log.Println(entry.Info())
	}

	return nil
}
