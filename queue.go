package address

import "fmt"

type queue []int

var last int

func (q *queue) Push(n int) {
	last = n
	*q = append(*q, n)
}

func (q *queue) Pop() (n int, e error) {
	if q.Len() <= 0 {
		e = fmt.Errorf("Index out of bounds.")
		return
	}

	n = (*q)[0]
	*q = (*q)[1:]
	return
}

func (q *queue) Peek() (n int) {
	n = (*q)[0]
	return
}
func (q *queue) PeekLast() (n int) {
	return last
}
func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) Get() []int {
	return *q
}
