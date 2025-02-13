package main

import (
	"fmt"
	"math"
)

func asteroidCollision(asteroids []int) []int {
	if len(asteroids) < 2 {
		return asteroids
	}
	res := []int{}
	for i := 0; i < len(asteroids); i++ {
		if len(res) == 0 { // Если текущий массив пуст - добавляем астероид, ведь сравнивать не с чем
			res = append(res, asteroids[i])
			continue
		}
		if res[i-1] < 0 { // Если астероди летит влево, то сравнивать с ним бесполезно
			res = append(res, asteroids[i])
		} else { // В ином случае, нам нужно сравнить астероид с каждым астероидом из масства в результатами
			for j := len(res) - 1; j >= 0; j-- {
				if asteroids[i] < 0 {
					if math.Abs(float64(res[j])) > math.Abs(float64(asteroids[i])) {
						break
					} else if math.Abs(float64(res[j])) == math.Abs(float64(asteroids[i])) {
						res = res[:j]
						break
					} else {
						res = res[:j]
						res = append(res, asteroids[i])
						continue
					}
				} else {
					res = append(res, asteroids[i])
					break
				}
			}
		}

	}
	return res
}

func main() {
	test := []int{10, 2, -5}
	fmt.Print(asteroidCollision(test))
}
