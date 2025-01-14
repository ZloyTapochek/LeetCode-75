package main

import "fmt"

func maxArea(height []int) int {

	length := len(height)
	current_square := 1
	if length == 2 {
		if height[0] < height[1] {
			return height[0] * 1
		} else {
			return height[1] * 1
		}
	}
	for left, right := 0, length-1; left != right; {
		if height[left] < height[right] {
			if height[left]*len(height[left:right]) >= current_square {
				current_square = height[left] * len(height[left:right])
			}
			left++
		} else {
			if height[right]*len(height[left:right]) >= current_square {
				current_square = height[right] * len(height[left:right])
			}
			right--
		}
	}
	return current_square

}
func main() {
	topics := []int{8, 7, 2, 1}
	fmt.Println(maxArea(topics))
}
