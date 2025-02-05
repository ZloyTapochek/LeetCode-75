package main

import "fmt"

func removeStars(s string) string {
	res := []rune{}

	for _, i := range s {
		if i == '*' {
			res = res[:len(res)]

		} else {
			res = append(res, i)
		}

	}
	return string(res)
}

func main() {
	fmt.Print(removeStars("leet**cod*e"))
}
