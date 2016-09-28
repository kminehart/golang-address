package address

import "fmt"

type queue []*string

func (q *queue) Push(n *string) {
	*q = append(*q, n)
}

func (q *queue) Pop() (n *string, e error) {
	if q.Len() <= 0 {
		e = fmt.Errorf("Index out of bounds.")
		return
	}

	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *queue) Peek() (n *string) {
	n = (*q)[0]
	return
}

func (q *queue) Len() int {
	return len(*q)
}
