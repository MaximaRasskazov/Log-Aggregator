package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"path/filepath" // Работа с путем к логам: Папка, Файлы
	"strings"
)

// сделать вывод в файл с указанным именем

// игнор регистра ключевых слов error ERROR

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

	type result struct {
		fileName string
		err      error
	}

	resultChan := make(chan result, len(files))

	// Чтение каждого файла из введеной директории
	for _, file := range files {
		go func(file os.DirEntry) {
			err := processFile(dir, file, keywords)
			resultChan <- result{
				fileName: file.Name(),
				err:      err,
			}
		}(file)
	}

	for i := 0; i < len(files); i++ {
		if res := <-resultChan; res.err != nil {
			log.Printf("Ошибка обработки файла %s: %v", res.fileName, res.err)
		}
	}

	close(resultChan)

	return nil
}

// Объявляем "Флаги" для CLI - командной строки,
// Далее "Парсим" ввод пользователя
func parseFlags() (string, []string) {
	dir := flag.String("dir", "test-logs", "Директория с логами")
	keywords := flag.String("keywords", "ERROR", "Ключевые слова для поиска")
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

func analyzeLine(line, fileName string, keywords []string, lineNumber int) {
	for _, kw := range keywords {
		keyword := strings.TrimSpace(kw)
		if keyword == "" {
			continue
		}

		keywordLower := strings.ToLower(keyword)
		lineLower := strings.ToLower(line)
		keywordLen := len(keywordLower)

		if keywordLen == 0 {
			continue
		}

		start := 0
		for {
			// Ищем вхождение в подстроке (начиная с позиции start)
			idx := strings.Index(lineLower[start:], keywordLower)
			if idx == -1 {
				break
			}

			// Реальная позиция в оригинальной строке
			realIdx := start + idx
			foundWord := line[realIdx : realIdx+keywordLen]
			keywordColor := colorForKeyword(keyword)

			log.Printf(
				"Найдено %s[%s]\033[0m в файле [%s]\033[0m на строке [%d]",
				keywordColor,
				foundWord,
				fileName,
				lineNumber,
			)

			// Перемещаем позицию поиска
			start = realIdx + keywordLen

			// Проверяем выход за границы строки
			if start >= len(lineLower) {
				break
			}
		}
	}
}

func colorForKeyword(key string) string {
	switch strings.ToUpper(key) {
	case "ERROR":
		return "\033[31m" // Красный
	case "WARNING":
		return "\033[33m" // Жёлтый
	case "INFO":
		return "\033[32m"
	default:
		return "\033[0m" // Сброс
	}
}
