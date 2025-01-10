package main

import (
	"fmt"
	"math"
)

func increasingTriplet(nums []int) bool {
	m1, m2 := math.MaxInt64, math.MaxInt64
	for _, i := range nums {
		if i <= m1 {
			m1 = i
		} else if i <= m2 {
			m2 = i
		} else {
			return true
		}
	}
	return false
}

func main() {
	res := []int{20, 100, 10, 12, 5, 13}
	fmt.Println(increasingTriplet(res))
}
