package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	res := make([]bool, 0, len(candies))
	maxcand := 0
	for _, value := range candies {
		if value > maxcand {
			maxcand = value
		}
	}
	for i := range len(candies) {
		if candies[i]+extraCandies >= maxcand {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}
func main() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithCandies(candies, extraCandies)
	fmt.Println(result) // Вывод: [true true true false true]
}
