package main

import (
	"os"
	"testing"
)

func TestRun(t *testing.T) {
	testDir := os.Getenv("TEST_DIR")
	if testDir == "" {
		t.Fatal("TEST_DIR not set")
	}

	logChan := make(chan logEntry, 10)
	keywords := []string{"ERROR", "FATAL"}

	err := run(logChan, testDir, keywords)
	if err != nil {
		t.Fatalf("run() failed: %v", err)
	}

	// Проверяем количество найденных ошибок
	expectedErrors := 3
	if len(logChan) != expectedErrors {
		t.Errorf("Expected %d errors, found %d", expectedErrors, len(logChan))
	}
}
