package main

import (
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

// Парсинг Флагов командной строки
func parseFlags() (string, []string) {
	dir := flag.String("dir", "test-logs", "Директория с логами")
	keywords := flag.String("keywords", "ERROR, FAIL", "Ключевые слова для поиска")
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
	content, err := readFileContent(fullPath)
	if err != nil {
		return err
	}

	analyzeContent(content, file.Name(), keywords)

	return nil
}

// Чтение содержимого файла: достаем содержимое
func readFileContent(fullPath string) (string, error) {
	text, err := os.ReadFile(fullPath)
	if err != nil {
		return "", err
	}

	return string(text), nil
}

// Анализ содержимого файла: ищем ключевые слова в файле (логе)
func analyzeContent(content, filename string, keywords []string) {
	for _, keyword := range keywords {
		if strings.Contains(content, keyword) {
			log.Printf("Найдено [%s] в файле [%s]", keyword, filename)
		}
	}
}
