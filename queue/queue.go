package queue

type Linked [] int

func (q *Linked) Push(v int) {
	*q = append(*q, v)

}

func (q *Linked) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head

}

func (q *Linked) IsEmpty() int {
	return len(*q)
}
