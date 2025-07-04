package main

import (
	"flag"
	"log"
	"os" // Работа с путем к логам: Папка, Файлы
	"strings"
)

// сделать вывод в файл с указанным именем

func main() {
	dir, keywords := parseFlags()
	err := run(dir, keywords)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}
}

// Объявляем "Флаги" для CLI - командной строки,
// Далее "Парсим" ввод пользователя
func parseFlags() (string, []string) {
	dir := flag.String("dir", "test-logs", "Директория с логами")
	keywords := flag.String("keywords", "ERROR", "Ключевые слова для поиска")
	flag.Parse()

	return *dir, strings.Split(*keywords, ",")
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
