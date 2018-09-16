package _recover

import "fmt"

func TryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occurred.", err)
		} else {
			panic(r)
		}
	}()

	a := 0
	b := 5 / a
	fmt.Println(b)
}
