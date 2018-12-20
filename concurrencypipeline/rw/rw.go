package rw

import (
	"bufio"
	"fmt"
	"os"
	"polaris/concurrencypipeline/pipeline"
)

func PrintFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		if count == 100 {
			break
		}
		count++
	}
}
func WriteToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	pipeline.WriteSink(writer, p)

}
