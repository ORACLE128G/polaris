package main

import (
	"fmt"
	"unicode/utf8"
)

func lenz(s string) {
	for _, ch := range []byte(s) {
		fmt.Printf("%X ", ch)
	}

	fmt.Println()
	for _, ch := range s {
		fmt.Printf("%X ", ch)
	}
}

func utf(s string) {
	fmt.Println("Rune count:", utf8.RuneCountInString(s))
	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)

	}

}
func main() {
	lenz("1302岁的龙猫")
	utf("1302岁的龙猫")
}
