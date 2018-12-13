package main

import (
	"fmt"
	"strings"
)

func isSuffle(s1 string, s2 string) bool {
	n, m := len(s1), len(s2)
	if n != m {
		return false
	}
	for i := 0; i < n; i++ {
		if s1[i] == s2[i] {
			return false
		}
		if !strings.Contains(s2, string(s1[i])) {
			return false
		}
		if strings.Count(s1, string(s1[i])) != strings.Count(s2, string(s1[i])) {
			return false
		}
	}
	return true
}

func main() {
	s1 := "abcd1234"
	s2 := "abcd4321"
	fmt.Println(s1, s2, isSuffle(s1, s2))
	s1 = "abcd1234"
	s2 = "4321dcba"
	fmt.Println(s1, s2, isSuffle(s1, s2))
	s1 = "aab"
	s2 = "baa"
	fmt.Println(s1, s2, isSuffle(s1, s2))
}
