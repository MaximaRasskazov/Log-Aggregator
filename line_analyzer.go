package main

import (
	"log"
	"strings"
)

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
