package mock_stack_adv

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

type MyStack struct {
	stack []int
}

func Constructor() MyStack {
	return MyStack{}
}

func (s *MyStack) Push(x int) {
	s.stack = append(s.stack, x)
}

func (s *MyStack) Pop() int {
	if len(s.stack) == 0 {
		return 0
	}

	for i := 0; i < len(s.stack)-1; i++ {
		val := s.stack[0]
		s.stack = s.stack[1:]
		s.stack = append(s.stack, val)
	}

	val := s.stack[0]
	s.stack = s.stack[1:]
	return val
}

func (s *MyStack) Top() int {
	if len(s.stack) == 0 {
		return 0
	}

	for i := 0; i < len(s.stack)-1; i++ {
		val := s.stack[0]
		s.stack = s.stack[1:]
		s.stack = append(s.stack, val)
	}

	val := s.stack[0]
	s.stack = s.stack[1:]
	// 弹出之后需要再添加进去
	s.stack = append(s.stack, val)
	return val
}

func (s *MyStack) Empty() bool {
	return len(s.stack) == 0
}
