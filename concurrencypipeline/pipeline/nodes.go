package pipeline

import (
	"encoding/binary"
	"io"
	"log"
	"math/rand"
	"sort"
	"time"
)

var startTime time.Time

const blockSize = 2048

func Init() {
	startTime = time.Now()
}
func ArraySource(ints ... int) <-chan int {

	out := make(chan int, len(ints))

	go func() {
		for _, v := range ints {
			out <- v
		}
		close(out)
	}()
	return out
}

func InMemSort(in <-chan int) <-chan int {
	out := make(chan int, blockSize)
	go func() {
		var all []int
		for v := range in {
			all = append(all, v)
		}
		log.Printf("Read done: %v \n", time.Now().Sub(startTime))
		// Sort
		sort.Ints(all)
		log.Printf("InMemSort done: %v \n", time.Now().Sub(startTime))

		// Output
		for _, v := range all {
			out <- v
		}
		close(out)
	}()
	return out
}

func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, blockSize)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2

		for ok1 || ok2 {
			// v1 has value
			if !ok2 || (ok1 && v1 <= v2) {
				out <- v1
				// continue take out a next value
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
		log.Printf("Merge done: %v \n", time.Now().Sub(startTime))
	}()
	return out
}

func MergeN(inputs ...<-chan int) <-chan int {
	if len(inputs) == 1 {
		return inputs[0]
	}

	mid := len(inputs) / 2
	return Merge(
		MergeN(inputs[:mid]...),
		MergeN(inputs[mid:]...),
	)
}

func ReaderSource(reader io.Reader, chunkSize int) <-chan int {
	out := make(chan int, blockSize)
	go func() {
		// 1byte = 8bit
		// 1int = 8byte = 64bit
		b := make([] byte, 8)
		bytesRead := 0
		for {
			n, err := reader.Read(b)
			bytesRead += n
			if n > 0 {
				// transform b to Uint64
				v := int(binary.BigEndian.Uint64(b))
				out <- v
			}

			if err != nil ||
				(chunkSize != -1 && bytesRead >= chunkSize) {
				break
			}
		}
		close(out)
	}()
	return out
}

func WriteSink(write io.Writer, in <-chan int) {
	for v := range in {
		buf := make([]byte, 8)
		binary.BigEndian.PutUint64(buf, uint64(v))
		write.Write(buf)
	}
}

func RandomSource(count int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < count; i++ {
			out <- rand.Int()
		}
		close(out)
	}()
	return out
}
