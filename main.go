package main

import (
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
)

func dataProcessing(dirPath string, file os.DirEntry, check *bool) {
	// Собираем полный путь к файлу
	fullPath := filepath.Join(dirPath, file.Name())

	// Открываем файл
	f, err := os.Open(fullPath)
	if err != nil {
		log.Printf("Ошибка открытия файла %s: %v", fullPath, err)
		return
	}
	defer f.Close()

	// Читаем файл
	data := make([]byte, 100)
	count, err := f.Read(data)
	if err != nil && err != io.EOF {
		log.Printf("Ошибка чтения файла %s: %v", fullPath, err)
		return
	}

	log.Printf("Прочитано %d байт из %s: %q\n", count, file.Name(), data[:count])
	*check = true
}

func main() {
	// Задаем текстовые константы вместо JSON
	const (
		descDir       = "Путь к директории с логами"
		errorDir      = "Директория, которую вы ввели, не найдена"
		errorFiles    = "Файлы не были обнаружены"
		scanningFiles = "Сканируем файлы из папки"
	)

	// Обработка флагов
	dir := flag.String("dir", "test-logs", descDir)
	flag.Parse()

	// Проверка существования директории
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatalf("%s: %s", errorDir, *dir)
	}

	log.Printf("%s: %s", scanningFiles, *dir)

	// Чтение содержимого директории
	files, err := os.ReadDir(*dir)
	if err != nil {
		log.Fatalf("Ошибка чтения директории: %v", err)
	}

	check := false // Флаг обнаружения файлов

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		dataProcessing(*dir, file, &check)
	}

	if !check {
		log.Printf("%s", errorFiles)
	}
}
