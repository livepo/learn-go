package main

import "fmt"


// "pwwkew"
func Max(a, b int) int {
	if a < b {
		return b
	} else {
		return a
	}
}


func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]int)
	length := 0
	i, j := 0, 0

	for j < len(s) {
		if _, ok := m[s[j]]; ok {
			for i < j {
				if s[i] == s[j] {
					i ++
					j ++
					break
				} else {
					delete(m, s[i])
					i ++
				}
			}
		} else {
			m[s[j]] = 1
			length = Max(j-i+1, length)
			j ++
		}
	}
	return length
}


func main() {
	s := "pwwkew"
	// s = "aab"
	// s = "dvdf"
	fmt.Println(lengthOfLongestSubstring(s))
}
