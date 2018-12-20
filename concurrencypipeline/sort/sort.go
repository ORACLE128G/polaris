package main

import (
	"bufio"
	"fmt"
	"os"
	"polaris/concurrencypipeline/pipeline"
)

func main() {
	// Run ../main.go first.
	// Then to run here.
	p := createPipeline("large.in", 800000000, 4)
	writeToFile(p, "large.out")
	printFile("large.out")
}
func printFile(filename string) {
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
func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	pipeline.WriteSink(writer, p)

}
func createPipeline(fileName string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	// start count time cost
	pipeline.Init()
	var sortResult []<-chan int
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)
		sortResult = append(sortResult,
			pipeline.InMemSort(source))
	}

	return pipeline.MergeN(sortResult...)
}
