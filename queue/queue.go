package queue

type Linked [] interface{}

// Push to Queue.
func (q *Linked) Push(v interface{}) {
	*q = append(*q, v)

}

// Popping Queue of pop.
func (q *Linked) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

// Queue is empty.
func (q *Linked) IsEmpty() bool {
	return len(*q) == 0
}
