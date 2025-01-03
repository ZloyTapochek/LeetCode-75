package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	tests := []struct {
		s        string
		t        string
		expected bool
	}{
		{"abc", "ahbgdc", true},  // "abc" является подпоследовательностью "ahbgdc"
		{"axc", "ahbgdc", false}, // "axc" не является подпоследовательностью "ahbgdc"
		{"", "ahbgdc", true},     // Пустая строка является подпоследовательностью любой строки
		{"abc", "", false},       // "abc" не является подпоследовательностью пустой строки
		{"abc", "abc", true},     // "abc" является подпоследовательностью "abc"
		{"a", "bba", false},      // "a" не является подпоследовательностью "bba"
	}

	for _, test := range tests {
		result := isSubsequence(test.s, test.t)
		if result != test.expected {
			t.Errorf("isSubsequence(%q, %q) = %v; expected %v", test.s, test.t, result, test.expected)
		}
	}
}
