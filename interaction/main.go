package main

import "fmt"

func interaction(in ...[2] int) [] int {
	var r = [] int{0, 0}
	for i := range in {
		t := in[i]
		if t[0] > t[1] {
			return []int{0, 0}
		}
		if t[0] < 0 || t[1] < 0 {
			return [] int{0, 0}
		}
		if r[0] < t[0] || r[0] == 0 {
			r[0] = t[0]
		}
		if r[1] > t[1] || r[1] == 0 {
			r[1] = t[1]
		}
	}
	if r[0] > r[1] {
		return [] int{0, 0}
	}
	return r

}
func main() {
	r := interaction(
		[2] int{1, 2},
		[2] int{2, 3},
		[2]int{1, 5},
		[2]int{0, -1})
	fmt.Println(r)
}
