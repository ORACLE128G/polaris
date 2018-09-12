package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strconv"
)

func eval(a, b, c int, d string) string {

	fmt.Println(
		a,
		b,
		c,
		d,
	)
	return strconv.FormatBool(false)
}

func div(a, b int) (q, r int) {

	return a % b, a / b
}

func err(a int) (string, error) {
	switch {
	case 1 == a:
		return "success", nil
	case 2 == a:
		fmt.Print("fail")
		return "success", nil
	default:
		return "success", fmt.Errorf("unsupported num: %s", "错误")
	}
}
func add(a, b int) string {
	return strconv.Itoa(a + b)

}
func apply(op func(int, int) string, a, b int) (string) {
	fmt.Printf("Calling %s with %d , %d args:",
		runtime.FuncForPC(reflect.ValueOf(op).Pointer()).Name(),
		a,
		b,
	)

	return op(a, b)

}

func sum(num ... int) int {
	s := 0
	for i := range num {

		s += num[i]
	}
	return s
}
func main() {

	fmt.Println(sum(1, 2, 3, 4, 5))
	i := apply(add, 1, 2)
	fmt.Println("返回值：" + i)
	if res, err := err(34); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
	q, _ := div(1, 2)
	fmt.Println(q)
	fmt.Println(eval(1, 2, 3, "4"))
}
