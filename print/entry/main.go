package main

import (
	"polaris/print"
	"strings"
)

func main() {
	// Use print/Print for implement io/Reader
	print.Print(strings.NewReader(`asd
asd
asd

asdasd`))
}
