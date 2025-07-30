package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

type logEntry struct {
	File    string
	Line    int
	Keyword string
	Message string
	Time    time.Time
}

func main() {
	logChan := make(chan logEntry, 1000)
	defer close(logChan)

	go logWriter(logChan)

	dir, keywords := parseFlags()

	log.Printf("Сканируем директорию: %s", dir)
	log.Printf("Ищем ключевые слова: %v", keywords)

	err := run(logChan, dir, keywords)
	if err != nil {
		log.Fatalf("Ошибка: %v", err)
	}

	time.Sleep(2 * time.Second)
	log.Println("Программа завершена")
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
func run(logChan chan<- logEntry, dir string, keywords []string) error {
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
			err := processFile(dir, file, keywords, logChan)
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
