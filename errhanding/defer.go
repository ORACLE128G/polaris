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

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}

}

func WriteFile(fileName string) {
	file, err := os.OpenFile(fileName, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(pathError)
		} else {
			fmt.Printf("%s %s %s \n",
				pathError.Op,
				pathError.Path,
				pathError.Err)
		}
	}
	// file, e := os.Create(fileName)
	defer file.Close()
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	fi := fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, fi())
	}
}
