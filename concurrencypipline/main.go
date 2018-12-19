package main

import (
	"fmt"
	"polaris/concurrencypipline/pipline"
)

func main() {
	out := pipline.Merge(
		pipline.InMemSort(pipline.ArraySource(1, 5, 2, 10, 6, 3)),
		pipline.InMemSort(pipline.ArraySource(3, 4, 9, 0, 11, 8)),
	)

	for v := range out {
		fmt.Println(v)
	}
}
