package main

func mergeAlternately(word1 string, word2 string) string {
	merged := ""
	len1 := len(word1)
	len2 := len(word2)
	if len1 > len2 {
		for i := 0; i < len1; i++ {
			merged += string(word1[i])
			if i < len2 {
				merged += string(word2[i])
			}
		}
	} else {
		for i := 0; i < len2; i++ {
			if i < len1 {
				merged += string(word1[i])
			}
			merged += string(word2[i])
		}
	}

	return merged
}
