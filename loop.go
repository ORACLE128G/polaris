package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func printFile(fileName string) string{
	file, err := os.Open(fileName)
	if err != nil {
		panic("Panic...")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return "success"
}
func forever (n int) {
	for ;n <= 100; n++ {
		fmt.Println("abc")
	}
}
func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func main() {

	forever(1)
	printFile("abc.txt")
	fmt.Println(
		convertToBin(5),
		convertToBin(13),

	)
}
