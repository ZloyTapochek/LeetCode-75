package main

import (
	"reflect"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	tests := []struct {
		candies      []int
		extraCandies int
		expected     []bool
	}{
		{[]int{2, 3, 5, 1, 3}, 3, []bool{true, true, true, false, true}},
		{[]int{4, 2, 1, 1, 2}, 1, []bool{true, false, false, false, false}},
		{[]int{12, 1, 12}, 10, []bool{true, false, true}},
	}

	for _, test := range tests {
		result := kidsWithCandies(test.candies, test.extraCandies)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("Для candies: %v и extraCandies: %d, ожидаемый результат: %v, полученный: %v", test.candies, test.extraCandies, test.expected, result)
		}
	}
}
