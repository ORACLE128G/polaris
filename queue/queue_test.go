package queue

import "fmt"

func ExampleLinked_Push() {
	q := Linked{1}
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)

	fmt.Println(q)
	// Output:
	// [1 2 3 4 5]
}
