package main

import "strings"

func colorForKeyword(key string) string {
	switch strings.ToUpper(key) {
	case "ERROR":
		return "\033[31m" // Красный
	case "WARNING":
		return "\033[33m" // Жёлтый
	case "INFO":
		return "\033[32m" // Зеленый
	default:
		return "\033[0m" // Сброс
	}
}
