package main

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func createTestFile(path string, mtime time.Time) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	f.Close()
	return os.Chtimes(path, time.Time{}, mtime)
}

func TestFindRecentFiles(t *testing.T) {
	tempDir := t.TempDir()

	if err := createTestFile(filepath.Join(tempDir, "file1.txt"), time.Now().AddDate(0, 0, -1)); err != nil {
		t.Fatal(err)
	}
	if err := createTestFile(filepath.Join(tempDir, "file2.txt"), time.Now().AddDate(0, 0, -10)); err != nil {
		t.Fatal(err)
	}

	files, err := findRecentFiles(tempDir, 7)
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 1 {
		t.Errorf("wanted 1 file got %d", len(files))
	}
	if files[0] != filepath.Join(tempDir, "file1.txt") {
		t.Errorf("wanted %s, got %s", filepath.Join(tempDir, "file1.txt"), files[0])
	}
}
