package main

import (
	"strings"
	"time"
)

func analyzeLine(line, fileName string, keywords []string, lineNumber int, logChan chan<- logEntry) {
	lineLower := strings.ToLower(line)

	for _, keyword := range keywords {
		keyword := strings.TrimSpace(keyword)
		if keyword == "" {
			continue
		}

		keywordLower := strings.ToLower(keyword)
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

			logChan <- logEntry{
				File:    fileName,
				Line:    lineNumber,
				Keyword: keyword,
				Message: foundWord,
				Time:    time.Now(),
			}

			time.Sleep(1 * time.Second)

			// Перемещаем позицию поиска
			start = realIdx + keywordLen

			// Проверяем выход за границы строки
			if start >= len(lineLower) {
				break
			}
		}
	}
}
