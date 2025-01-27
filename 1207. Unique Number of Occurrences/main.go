package main

func uniqueOccurrences(arr []int) bool {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		_, key := m[arr[i]]
		if key {
			m[arr[i]]++
		} else {
			m[arr[i]] = 1
		}
	}
	occurrencesMap := make(map[int]bool)
	for _, val := range m {
		_, exist := occurrencesMap[val]
		if exist {
			return false
		}
		occurrencesMap[val] = true
	}
	return true
}
