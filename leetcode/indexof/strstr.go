package main

import "fmt"

func main() {
	fmt.Println(strStr("aa", "a"))
}

// Str str, IndexOf
// https://leetcode.com/problems/implement-strstr/submissions/
func strStr(haystack string, needle string) int {
	if len(needle) > len(haystack) {
		return -1
	}
	if haystack == needle || needle == "" {
		return 0
	}
	var m, n = len(haystack), len(needle)
	var final, pointer, count = 0, 0, 0
	for pointer < m {
		if haystack[pointer] == needle[count] {
			if count == n -1 {
				return final
			}
			count++
			pointer++
		} else {
			final++
			pointer = final
			count = 0
		}
	}
	return -1
}
