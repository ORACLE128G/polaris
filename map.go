package main

import "fmt"

func maps() {
	m := map[string]string{
		"key": "value",
	}
	fmt.Println(m)

	m1 := make(map[string]int8)
	fmt.Println(m1)
}

func forMaps() {
	m := map[string]string{
		"key":  "value",
		"key1": "value",
		"key2": "value",
	}
	for k, v := range m {
		fmt.Println(k, v)

		if s, ok := m[k]; ok {
			fmt.Println(s, ok)
			delete(m, k)
		}
	}
	fmt.Println(m)
}

func main() {

	maps()
	forMaps()
}
