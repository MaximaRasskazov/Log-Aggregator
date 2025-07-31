package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

func logWriter(logChan <-chan logEntry) {
	var buffer []logEntry
	const batch = 1000
	ticker := time.NewTicker(5 * time.Second) // Таймер сброса буфера

	for {
		select {
		case entry, ok := <-logChan:
			if !ok {
				return
			}

			buffer = append(buffer, entry)

			if len(buffer) >= batch {
				writeBatch(buffer)
				buffer = buffer[:0]
			}

		case <-ticker.C:
			if len(buffer) > 0 {
				writeBatch(buffer)
				buffer = buffer[:0]
			}
		}

	}
}

func writeBatch(entries []logEntry) {
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Time.Before(entries[j].Time)
	})

	file, err := os.OpenFile("output.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Ошибка записи лога %v\n", err)
		return
	}

	defer file.Close()

	for _, entry := range entries {
		line := fmt.Sprintf(
			"Найдено [%s] в файле [%s] на строке [%d]\n",
			entry.Message,
			entry.File,
			entry.Line,
		)

		if _, err := file.WriteString(line); err != nil {
			fmt.Printf("Ошибка записи строки: %v\n", err)
		}
	}
}
