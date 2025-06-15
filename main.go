package main

/*
1. Обработка ошибок
2. Логика чтения и вывода
*/

import (
	"encoding/json"
	"flag"
	"log"
	"os"
)

type Config struct {
	DescDir       string `json:"descDir"`
	ErrorDir      string `json:"errorDir"`
	ErrorFiles    string `json:"errorFiles"`
	ScanningFiles string `json:"scanningFiles"`
}

// Загрузка данных из json файла, для дальнейшего вывода
func loadConfig(filename string) (Config, error) {
	var config Config

	file, _ := os.Open(filename)
	defer file.Close()

	decoder := json.NewDecoder(file)
	err := decoder.Decode(&config)

	return config, err

}

func dataProcessing(file os.DirEntry, check *bool) {
	log.Print(file)
	*check = true
}

func main() {
	config, _ := loadConfig("archive.json")
	dir := flag.String("dir", "test-logs", config.DescDir)
	flag.Parse()

	// Проверка на чтение и наличие
	if _, err := os.Stat(*dir); os.IsNotExist(err) {
		log.Fatalf("%s %s", config.ErrorDir, *dir)
	}

	log.Printf("%s %s", config.ScanningFiles, *dir)

	files, _ := os.ReadDir(*dir)
	check := false

	for _, file := range files {
		// Пропускаем иттерацию, если это не файл, а директория
		if file.IsDir() {
			continue
		}

		dataProcessing(file, &check)
	}

	if !check {
		log.Printf("%s", config.ErrorFiles)
	}
}
