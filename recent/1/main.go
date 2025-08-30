package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func main() {
	dir := "."
	t := time.Now().AddDate(0, 0, -7)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.ModTime().After(t) {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
