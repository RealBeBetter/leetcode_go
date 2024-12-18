package mock_stack

// 225. 用队列实现栈
// https://leetcode.cn/problems/implement-stack-using-queues

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

type MyQueue []int

func (q *MyQueue) Push(x int) {
	*q = append(*q, x)
}

func (q *MyQueue) Pop() int {
	if len(*q) == 0 {
		return 0
	}

	popVal := (*q)[0]
	*q = (*q)[1:]
	return popVal
}

func (q *MyQueue) Top() int {
	if len(*q) == 0 {
		return 0
	}

	return (*q)[0]
}

func (q *MyQueue) Empty() bool {
	return len(*q) == 0
}

type MyStack struct {
	inQueue  MyQueue
	outQueue MyQueue
}

func Constructor() MyStack {
	return MyStack{
		inQueue:  MyQueue{},
		outQueue: MyQueue{},
	}
}

func (s *MyStack) Push(x int) {
	s.inQueue.Push(x)
}

func (s *MyStack) Pop() int {
	if len(s.inQueue) == 0 {
		return s.inQueue.Pop()
	}

	for len(s.inQueue) > 1 {
		s.outQueue.Push(s.inQueue.Pop())
	}

	pop := s.inQueue.Pop()

	s.inQueue, s.outQueue = s.outQueue, s.inQueue
	return pop
}

func (s *MyStack) Top() int {
	if len(s.inQueue) == 0 {
		return s.inQueue.Top()
	}

	for len(s.inQueue) > 1 {
		s.outQueue.Push(s.inQueue.Pop())
	}

	topVal := s.inQueue.Top()

	s.outQueue.Push(s.inQueue.Pop())
	s.inQueue, s.outQueue = s.outQueue, s.inQueue
	return topVal
}

func (s *MyStack) Empty() bool {
	return s.inQueue.Empty()
}
