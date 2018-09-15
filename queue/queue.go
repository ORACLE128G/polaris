package queue

type Linked [] interface{}

func (q *Linked) Push(v interface{}) {
	*q = append(*q, v)

}

func (q *Linked) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

func (q *Linked) IsEmpty() bool {
	return len(*q) == 0
}
