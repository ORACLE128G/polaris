package errhanding

import (
	"bufio"
	"fmt"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
	panic("Error occurred.")
}

func fibonacci() func() int{
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}

}

func WriteFile (fileName string) {
	file, e := os.Create(fileName)
	if e != nil{
		panic("Error occurred.")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fi := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fi())
	}
}
