package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)

	// Проход 1: вычисляем произведение всех элементов слева от каждого элемента
	leftProduct := 1
	for i := 0; i < n; i++ {
		answer[i] = leftProduct
		leftProduct *= nums[i]
	}

	// Проход 2: вычисляем произведение всех элементов справа от каждого элемента
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= rightProduct
		rightProduct *= nums[i]
	}

	return answer
}

func main() {
	var nums [4]int = [4]int{1, 2, 3, 4}
	result := productExceptSelf(nums[:])
	fmt.Println(result)
}
