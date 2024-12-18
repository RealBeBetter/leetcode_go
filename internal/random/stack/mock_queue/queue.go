package mock_queue

// 232. 用栈实现队列
// https://leetcode.cn/problems/implement-queue-using-stacks

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

type MyStack []int

func (s *MyStack) Push(x int) {
	*s = append(*s, x)
}

func (s *MyStack) Pop() int {
	if len(*s) == 0 {
		return 0
	}

	popVal := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return popVal
}

func (s *MyStack) Peek() int {
	if len(*s) == 0 {
		return 0
	}

	popVal := (*s)[len(*s)-1]
	return popVal
}

func (s *MyStack) Empty() bool {
	return len(*s) == 0
}

type MyQueue struct {
	inStack  MyStack
	outStack MyStack
}

func Constructor() MyQueue {
	return MyQueue{
		// 切片零值可用
		inStack:  MyStack{},
		outStack: MyStack{},
	}
}

func (q *MyQueue) Push(x int) {
	q.inStack.Push(x)
}

func (q *MyQueue) Pop() int {
	if q.outStack.Empty() {
		for !q.inStack.Empty() {
			q.outStack.Push(q.inStack.Pop())
		}
	}

	return q.outStack.Pop()
}

func (q *MyQueue) Peek() int {
	if q.outStack.Empty() {
		for !q.inStack.Empty() {
			q.outStack.Push(q.inStack.Pop())
		}
	}

	return q.outStack.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.inStack.Empty() && q.outStack.Empty()
}
