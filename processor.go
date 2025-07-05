package main

import (
	"bufio"
	"os"
	"path/filepath" // Работа с путем к логам: Папка, Файлы
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
func processFile(dir string, file os.DirEntry, keywords []string, logChan chan<- logEntry) error {
	if file.IsDir() {
		return nil
	}

	fullPath := filepath.Join(dir, file.Name())
	openedFile, err := os.Open(fullPath)
	if err != nil {
		return err
	}
	defer openedFile.Close()

	readFileContent(openedFile, keywords, file.Name(), logChan)

	return nil
}

func readFileContent(file *os.File, keywords []string, filename string, logChan chan<- logEntry) {
	scanner := bufio.NewScanner(file)
	lineNumber := 0

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		analyzeLine(line, filename, keywords, lineNumber, logChan)
	}

	// if err := scanner.Err(); err != nil {
	// 	fmt.Errorf("ошибка сканирования %s: %w", filename, err)
	// }
}
