package main

import (
	"bufio"
	"os"
	"path/filepath"
)

// Сканирование директории,
// которую укзывают в командой строке
func scanDirectory(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// Обработка одного файла
func processFile(dir string, file os.DirEntry, keywords []string) error {
	if file.IsDir() {
		return nil
	}

	fullPath := filepath.Join(dir, file.Name())
	openedFile, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer openedFile.Close()

	readFileContent(openedFile, keywords, file.Name())

	return nil
}

func readFileContent(file *os.File, keywords []string, filename string) {
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		analyzeLine(line, filename, keywords, lineNumber)
	}

	// if err := scanner.Err(); err != nil {
	// 	fmt.Errorf("ошибка сканирования %s: %w", filename, err)
	// }
}
