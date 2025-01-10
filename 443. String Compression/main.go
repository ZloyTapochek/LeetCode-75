package main

import "fmt"

func compress(chars []byte) int {
	length := len(chars)
	if length == 1 {
		return 1
	}
	res := []byte{}
	current := chars[0]
	count := 0
	for j, i := 0, 0; i < length; j, i = j+1, i+1 {
		if i == length-1 {
			res = append(res, current)
			res = append(res, []byte(fmt.Sprintf("%d", count))...)
			break
		}

		if chars[i] == current {
			count++
		} else {
			res = append(res, current)
			res = append(res, []byte(fmt.Sprintf("%d", count))...)
			count = 1
			current = chars[i]
		}

	}
	fmt.Println(res)
	return len(res)
}
func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	compressed := compress(chars)
	fmt.Println(compressed) // Вывод: [a 2 b 2 c 3]
	chars1 := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	compressed1 := compress(chars1)
	fmt.Println(compressed1) // Вывод: [a 2 b 2 c 3]
}
