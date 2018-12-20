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

	// Before enable the buffer for Reader

	//2018/12/20 22:11:22 Read done: 9.3867035s
	//2018/12/20 22:11:22 Read done: 9.428592s
	//2018/12/20 22:11:22 Read done: 9.4944156s
	//2018/12/20 22:11:22 Read done: 9.7008637s
	//2018/12/20 22:11:34 InMemSort done: 20.8224555s
	//2018/12/20 22:11:34 InMemSort done: 20.8628778s
	//2018/12/20 22:11:34 InMemSort done: 20.8718539s
	//2018/12/20 22:11:34 InMemSort done: 21.0598563s
	//2018/12/20 22:13:59 Merge done: 2m46.6034671s
	//2018/12/20 22:13:59 Merge done: 2m46.6034671s
	//2018/12/20 22:13:59 Merge done: 2m46.6034671s

	// The buffer enabling

	//2018/12/20 22:20:28 Read done: 3.5886295s
	//2018/12/20 22:20:28 Read done: 3.9077763s
	//2018/12/20 22:20:28 Read done: 3.9307151s
	//2018/12/20 22:20:28 Read done: 3.9576432s
	//2018/12/20 22:20:39 InMemSort done: 15.0989995s
	//2018/12/20 22:20:40 InMemSort done: 15.360847s
	//2018/12/20 22:20:40 InMemSort done: 15.3778013s
	//2018/12/20 22:20:40 InMemSort done: 15.4416307s
	//2018/12/20 22:21:03 Merge done: 38.3544587s
	//2018/12/20 22:21:03 Merge done: 38.3544587s
	//2018/12/20 22:21:03 Merge done: 38.3554571s
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
