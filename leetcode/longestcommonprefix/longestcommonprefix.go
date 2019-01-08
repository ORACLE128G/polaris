package longestcommonprefix

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flow", "flower", "flo"}))
}

// Longest common prefix.
// see: https://leetcode.com/problems/longest-common-prefix/
func longestCommonPrefix(strs []string) (result string) {
	shortest := -1
	for _, s := range strs {
		if len(s) < shortest || shortest == -1 {
			shortest = len(s)
		}
	}
	//
	for i := 0; i < shortest; i++ {
		for _, s := range strs {
			// compare each character
			if s[i] != strs[0][i] {
				return
			}
		}
		result += string(strs[0][i])
	}
	return
}
