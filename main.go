package main

import (
	"fmt"
	"io/ioutil"
	"math/cmplx"
)

var (
	t = "3"
	f = "4"
)

const (
	g = "adsa"
	h = "中未能"
)

func variable() {
	var a int
	var b string
	fmt.Printf("%d %q\n", a, b)
}
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}
func variableTypeDeduction() {

	var a, b, c, s = 3, 4, 5, true
	fmt.Println(a, b, c, s)
}

func variableTypeDeduction1() {
	a, b, c, s := 3, 4, 5, true
	b = 5
	fmt.Println(a, b, c, s, t, f, g, h)
}
func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

const pass int = 0

func consts() {
	const name string = "asdasd"
	fmt.Println(name, pass)
}

func enums() {
	const (
		cpp = iota
		_
		golang
	)
	fmt.Println(cpp, golang)
}

const (
	b = 1 << (10 * iota)
	d
	e
)

func foreach() {

	fmt.Println()
}

func move() {
	a := 1 << 10
	fmt.Println(a)
}

func condition(v int) {
	if v > 1 {
		fmt.Println(v)
	} else if v > 2 {
		fmt.Println(v)
	} else {
		fmt.Println(v)
	}
}

func readFile() {
	const file = "abc.txt"
	if contents, err := ioutil.ReadFile(file); err != nil {
		fmt.Println("未读取到文件")
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func switchCase(cases int) {
	switch cases {
	case 1:
		fmt.Println("Hello")
		fallthrough
	case 2:
		fmt.Print("World.")
	default:
		fmt.Println("撒大苏打")
	}
}

func grade (score int) string {
	switch score {
	case -1:
		panic("Wrong score")
	case 10:
		fmt.Println("Low")
	case 20:
		fmt.Println("Too Low")
	default:
		fmt.Println("Good.")
	}
	return string(score)
}
func main() {
	/*fmt.Println("Hello World!")
	variable()
	variableInitialValue()
	variableTypeDeduction()
	variableTypeDeduction1()
	euler()
	consts()*/
	/*enums()
	move()
	fmt.Println(b, d, e)
	condition(2)
	readFile()*/

	// readFile()
	// switchCase(1)

	fmt.Println(grade(-1))
}
