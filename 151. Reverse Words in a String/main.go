package main

import (
	"slices"
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return ""
	}
	words := strings.Fields(strings.TrimSpace(s))
	//Создаем массив, чтобы потом его развернуть
	slices.Reverse(words) //Разворачиваем массив
	return strings.Join((words), " ")
}

func main() {

	reverseWords("the sky is blue")
	reverseWords("  hello world  ")
	reverseWords("a good   example")
}
