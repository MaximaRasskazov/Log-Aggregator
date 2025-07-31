package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	exitCode := m.Run()
	teardown()

	os.Exit(exitCode)
}

func setup() {
	testDir, _ := os.MkdirTemp("", "log-aggregator-test")
	os.Setenv("TEST_DIR", testDir)

	createTestFiles(testDir)
}

func teardown() {
	testDir := os.Getenv("TEST_DIR")
	os.RemoveAll(testDir)
}

func createTestFiles(dir string) {
	// Файл 1
	file1 := filepath.Join(dir, "app.log")
	content1 := `ERROR: Database connection failed
INFO: Server started on port 8080
WARNING: High memory usage`
	os.WriteFile(file1, []byte(content1), 0644)

	// Файл 2
	file2 := filepath.Join(dir, "error.log")
	content2 := `FATAL: Critical system failure
ERROR: File not found`
	os.WriteFile(file2, []byte(content2), 0644)
}
