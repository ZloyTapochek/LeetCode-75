package main

func isSubsequence(s string, t string) bool {
	i, j, res := 0, 0, len(s)
	for i < len(t) && j < res {
		if t[i] == s[j] {
			j++
		}
		i++
	}
	return j == res
}
