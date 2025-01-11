package main

func moveZeroes(nums []int) {
	length := len(nums)
	// if len()

	for i, j := 0, length-1; i != j; i++ {
		if nums[i] == 0 {
			save := nums[i]
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, save)
		}

	}
}
