package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func findRecentFiles(root string, days int) ([]string, error) {
	var files []string

	cutoff := time.Now().AddDate(0, 0, -days)

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().After(cutoff) {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func main() {
	root := "."
	days := 7

	files, err := findRecentFiles(root, days)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
	fmt.Println(files)
}
