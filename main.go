package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
)

/*
A struct to hold a key value pair corresponding to single member of a
map[string]uint32. This will be sued to sort the map data by value.
*/
type pair struct {
	name string
	size uint32
}

func main() {
	args := os.Args
	if len(args) != 2 {
		msg := fmt.Sprintf("Usage: %v <path_to_directory>", args[0])
		log.Fatal(msg)
	}

	fullPath := args[1]
	dirMap := make(map[string]uint32)

	// Recursively visit all filesystem entries at the provided path
	filepath.WalkDir(fullPath, func(path string, entry fs.DirEntry, err error) error {
		return walk(path, entry, err, dirMap)
	})

	// Declare a slice of pairs
	var pairs []pair

	// Add map entries to pairs slice
	for k, v := range dirMap {
		pairs = append(pairs, pair{k, v})
	}

	// Sort descending
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].size > pairs[b].size
	})

	// If the provided path had fewer than 10 entries, print all entries and their sizes in bytes
	if len(pairs) < 10 {
		for _, file := range pairs {
			fmt.Printf("%s, %d\n", file.name, file.size)
		}
	} else { // print only the largest 10 entries and their sizes in bytes
		for i := 0; i < 10; i++ {
			fmt.Printf("%s:  %d\n", pairs[i].name, pairs[i].size)
		}
	}

}

// Implement a WalkDirFunc that accepts a reference to a map[string]uint32
func walk(path string, entry fs.DirEntry, err error, dirMap map[string]uint32) error {
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
		dirMap[path] = uint32(size)
	}

	return nil
}
