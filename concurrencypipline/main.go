package main

import (
	"fmt"
	"polaris/concurrencypipline/pipline"
)

func main() {
	out := pipline.ArraySource(1, 5, 2, 10, 6, 3)
	for v := range out {
		fmt.Println(v)
	}
}
