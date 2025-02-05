package main

import "sort"

func closeStrings(word1 string, word2 string) bool {
	length_firstl := len(word1)
	length_secondl := len(word2)
	// Первый шаг: проверяем длину строк
	if length_firstl != length_secondl {
		return false
	}
	map_firstl := make(map[rune]bool)
	map_secondl := make(map[rune]bool)
	for _, i := range word1 {
		map_firstl[i] = true
	}
	for _, i := range word2 {
		map_secondl[i] = true
	}

	return CompareMaps(map_firstl, map_secondl)
}

func CompareMaps(map1, map2 map[rune]bool) bool {
	// Сравниваем длины мап
	if len(map1) != len(map2) {
		return false
	}

	// Проверяем, содержатся ли все ключи и значения из map1 в map2
	for key, value := range map1 {
		if val, exists := map2[key]; !exists || val != value {
			return false
		}
	}

	return true
}

// Решение:
func closeStrings(word1 string, word2 string) bool {
	counter := func(word string) (keys, vals [26]int) {
		for i := range word {
			keys[word[i]-'a'] = 1
			vals[word[i]-'a'] += 1
		}
		sort.Ints(vals[:])
		return keys, vals
	}
	keys1, vals1 := counter(word1)
	keys2, vals2 := counter(word2)
	return keys1 == keys2 && vals1 == vals2
}
