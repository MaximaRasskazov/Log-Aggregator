package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath" // Работа с путем к логам: Папка, Файлы
	"strings"
)

func main() {
	dir, keywords := parseFlags()
	err := run(dir, keywords)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}

// Основная логика программы
func run(dir string, keywords []string) error {
	files, err := scanDirectory(dir)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := processFile(dir, file, keywords)
		if err != nil {
			log.Printf("Ошибка обработки файла %s: %v", file.Name(), err)
		}
	}

	return nil
}

// Объявляем "Флаги" дял CLI - командной строки,
// Далее "Парсим" ввод пользователя
func parseFlags() (string, []string) {
	dir := flag.String("dir", "test-logs", "Директория с логами")
	keywords := flag.String(
		"keywords",
		"ERROR",
		"Ключевые слова для поиска",
	)
	flag.Parse()

	return *dir, strings.Split(*keywords, ",")
}

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

	readFileContent(openedFile, keywords, file.Name())

	return nil
}

func readFileContent(file *os.File, keywords []string, filename string) {
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	flag := 1

	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()
		analyzeLine(line, filename, keywords, lineNumber, &flag)
	}
}

func analyzeLine(line, filename string, keywords []string, lineNumber int, f *int) {
	for _, kw := range keywords {
		keyword := strings.TrimSpace(kw)
		if strings.Contains(line, keyword) {
			var filenameColor string
			keywordColor := colorForKeyword(keyword)

			if *f == 1 {
				filenameColor = "\033[32m"
			} else {
				filenameColor = "\033[0m"
			}

			log.Printf(
				"Найдено %s[%s]\033[0m в файле %s[%s]\033[0m на строке [%d]",
				keywordColor,
				keyword,
				filenameColor,
				filename,
				lineNumber,
			)

			*f = 0
		}
	}
}

func colorForKeyword(key string) string {
	switch key {
	case "ERROR":
		return "\033[31m" // Красный
	case "WARN":
		return "\033[33m" // Жёлтый
	default:
		return "\033[0m" // Сброс
	}
}
