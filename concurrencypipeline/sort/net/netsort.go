package main

import (
	"bufio"
	"os"
	"polaris/concurrencypipeline/pipeline"
	"polaris/concurrencypipeline/rw"
	"strconv"
)

func main () {
	p := createNetworkPipeline("small.in", 512, 4)
	rw.WriteToFile(p, "small.out")
	rw.PrintFile("small.out")
}


func createNetworkPipeline(fileName string,
	fileSize, chunkCount int) <-chan int {
	chunkSize := fileSize / chunkCount
	// start count time cost
	pipeline.Init()

	var sortAddr []string
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)
		source := pipeline.ReaderSource(bufio.NewReader(file), chunkSize)

		addr := ":" + strconv.Itoa(7000+i)
		pipeline.NetworkSink(addr,
			pipeline.InMemSort(source))
		sortAddr = append(sortAddr, addr)
	}

	var sortResults []<-chan int
	for _, v := range sortAddr {
		sortResults = append(sortResults, pipeline.NetworkSource(v))
	}
	return pipeline.MergeN(sortResults...)
}