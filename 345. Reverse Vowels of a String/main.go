package main

import (
	"strings"
)

func reverseVowels(s string) string {
	// Создаем строку со всеми гласными
	vowels := "aeiouAEIOU"
	dataLength := len(s)

	// Создаем срез для гласных и их индексов
	var writtenVowels []rune // Срез для гласных
	var indexes []int        // Срез для индексов

	// Проходим по строке и собираем гласные и их индексы
	for i := 0; i < dataLength; i++ {
		if strings.ContainsRune(vowels, rune(s[i])) { // Используем strings.ContainsRune для проверки
			writtenVowels = append(writtenVowels, rune(s[i]))
			indexes = append(indexes, i)
		}
	}

	// Переворачиваем собранные гласные
	for i, j := 0, len(writtenVowels)-1; i < j; i, j = i+1, j-1 {
		writtenVowels[i], writtenVowels[j] = writtenVowels[j], writtenVowels[i]
	}

	// Заменяем старые гласные новыми
	res := []rune(s) // Преобразуем строку в срез рун
	for i, j := 0, 0; i < dataLength; i++ {
		if j < len(indexes) && i == indexes[j] { // Проверяем, не вышли ли мы за пределы индексов
			res[i] = writtenVowels[j]
			j++
		}
	}
	return string(res)
}
