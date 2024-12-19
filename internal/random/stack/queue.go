package stack

type Queue []int

func (q *Queue) Push(x int) {
	*q = append(*q, x)
}

func (q *Queue) Pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func (q *Queue) Peek() int {
	return (*q)[0]
}

func (q *Queue) Empty() bool {
	return len(*q) == 0
}
