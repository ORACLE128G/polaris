package print

import (
	"bufio"
	"fmt"
	"io"
)

func Print(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
