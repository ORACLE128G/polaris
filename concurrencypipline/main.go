package main

import (
	"bufio"
	"fmt"
	"os"
	"polaris/concurrencypipline/pipline"
)

const (
	in   = "large.in"
	size = 100000000
)

func main() {
	file, err := os.Create(in)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	p := pipline.RandomSource(size)

	writer := bufio.NewWriter(file)
	pipline.WriteSink(writer, p)
	writer.Flush()

	file, err = os.Open(in)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipline.ReaderSource(bufio.NewReader(file), -1)

	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count == 100 {
			break
		}
	}
}

func Merge() {
	out := pipline.Merge(
		pipline.InMemSort(pipline.ArraySource(1, 5, 2, 10, 6, 3)),
		pipline.InMemSort(pipline.ArraySource(3, 4, 9, 0, 11, 8)),
	)

	for v := range out {
		fmt.Println(v)
	}
}
