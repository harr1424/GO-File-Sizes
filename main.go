package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"
)

/*
A struct to hold a key value EntryInfo corresponding to single member of a
map[string]uint32. This will be sued to sort the map data by value.
*/
type EntryInfo struct {
	name string
	size uint32
}

const NUM_ENTRIES = 10

// Implement a WalkDirFunc that accepts a reference to a slice pf EntryInfo
func walk(path string, entry fs.DirEntry, err error, pairs *[]EntryInfo) error {
	if err != nil {
		return err
	}

	if !entry.IsDir() {
		// add entry and it's size to dirMap
		file, err := os.Stat(path)
		if err != nil {
			log.Println("Error determining file size:", err)
		}

		newPair := EntryInfo{path, uint32(file.Size())}
		*pairs = append(*pairs, newPair)
	}

	return nil
}

func main() {
	// string representing directory to analyze
	var fullPath = ""

	args := os.Args
	if len(args) > 2 {
		msg := fmt.Sprintf("Usage: %v <path_to_directory>", args[0])
		log.Fatal(msg)
	} else if len(args) == 1 {
		cwd, err := os.Executable()
		if err != nil {
			log.Fatal("Unable to infer current working directory, try providing a path as program argument")
		}
		fullPath = filepath.Dir(cwd)
		log.Println("No path specified, defaulting to current working directory")
	} else {
		fullPath = args[1]
	}

	fmt.Printf("Searching for %v largest files in %v\n\n", NUM_ENTRIES, fullPath)

	programStart := time.Now()

	// Declare a slice of pairs to hold entry information
	var pairs []EntryInfo

	// Recursively visit all filesystem entries at the provided path
	filepath.WalkDir(fullPath, func(path string, entry fs.DirEntry, err error) error {
		return walk(path, entry, err, &pairs)
	})

	// Sort descending
	sort.Slice(pairs, func(a, b int) bool {
		return pairs[a].size > pairs[b].size
	})

	programDuration := time.Since(programStart).Seconds()

	// If the provided path had fewer than NUM_ENTRIES, print all entries and their sizes in bytes
	if len(pairs) < NUM_ENTRIES {
		for _, file := range pairs {
			fmt.Printf("%s, %d\n", file.name, file.size)
		}
	} else { // print only the largest NUM_ENTRIES and their sizes in bytes
		for i := 0; i < NUM_ENTRIES; i++ {
			fmt.Printf("%s:  %d\n", pairs[i].name, pairs[i].size)
		}
	}

	fmt.Printf("\nProgram completed in %v seconds.\n", programDuration)
}

// program completes in 0.26 seconds on average
