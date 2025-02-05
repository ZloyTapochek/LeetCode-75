package main

import "fmt"

func equalPairs(grid [][]int) int {
	// Объявляем переменные
	CountOfEqualPars := 0
	MapOfPars := make(map[int]int)

	// Пройдемся по строчкам
	for _, i := range grid {
		sum := 0
		for _, value := range i {
			sum += value
		}
		MapOfPars[sum]++

		if MapOfPars[sum] > CountOfEqualPars {
			CountOfEqualPars = MapOfPars[sum]
		}
	}
	// Пройдемся по столбикам

	for i := 0; i < len(grid); i++ {
		sum := 0
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
		}
		MapOfPars[sum]++

		if MapOfPars[sum] > CountOfEqualPars {
			CountOfEqualPars = MapOfPars[sum]
		}
	}

	return CountOfEqualPars
}
func main() {
	grid := [][]int{
		{3, 2, 1},
		{1, 7, 6},
		{2, 7, 7},
	}
	fmt.Print(equalPairs(grid))
}
