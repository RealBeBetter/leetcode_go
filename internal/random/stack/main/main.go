package main

import (
	"sort"
	"strconv"
)

// 20. 有效的括号
// https://leetcode.cn/problems/valid-parentheses
func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}

		prev := stack[len(stack)-1]
		if (prev == '(' && b == ')') || (prev == '[' && b == ']') ||
			(prev == '{' && b == '}') {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, b)
	}

	return len(stack) == 0
}

// 1047. 删除字符串中的所有相邻重复项
// https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string/
func removeDuplicates(s string) string {
	// 使用栈来保存已遍历过的数据，可以模拟出整个遍历过程
	stack := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		if len(stack) == 0 {
			stack = append(stack, b)
			continue
		}

		prev := stack[len(stack)-1]
		if prev == b {
			stack = stack[:len(stack)-1]
			continue
		}

		stack = append(stack, b)
	}

	return string(stack)
}

// 150. 逆波兰表达式求值
// https://leetcode.cn/problems/evaluate-reverse-polish-notation
func evalRPN(tokens []string) int {
	// 使用操作数栈进行操作即可
	numbers := make([]int, 0, len(tokens))
	for _, token := range tokens {
		if !(token == "+" || token == "-" || token == "*" || token == "/") {
			number, _ := strconv.Atoi(token)
			numbers = append(numbers, number)
			continue
		}

		// 遇到操作符开始计算，将结果入栈
		numbers = calc(numbers, token)
	}

	return numbers[0]
}

func calc(numbers []int, token string) []int {
	// 这里不考虑异常情况
	prev := numbers[len(numbers)-2]
	next := numbers[len(numbers)-1]
	numbers = numbers[:len(numbers)-2]

	var result int
	switch token {
	case "+":
		result = prev + next
	case "-":
		result = prev - next
	case "*":
		result = prev * next
	case "/":
		result = prev / next
	default:
	}

	numbers = append(numbers, result)
	return numbers
}

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum
func maxSlidingWindow(nums []int, k int) []int {
	ans := make([]int, len(nums)-k+1)
	idx := 0

	maxVal, maxIdx := calcMax(nums, 0, k-1)
	ans[idx] = maxVal
	idx++

	// 使用遍历的方式解决，用有限的临时变量存储上一个最值
	for i, j := 1, k; j < len(nums); i, j = i+1, j+1 {
		if maxIdx == i-1 {
			maxVal, maxIdx = calcMax(nums, i, j)
		} else {
			if nums[j] >= maxVal {
				maxVal = nums[j]
				maxIdx = j
			}
		}

		ans[idx] = maxVal
		idx++
	}

	return ans
}

func calcMax(nums []int, i, j int) (maxVal, maxIdx int) {
	maxIdx = i
	maxVal = nums[i]
	for i <= j {
		if nums[i] >= maxVal {
			maxVal = nums[i]
			maxIdx = i
		}
		i++
	}

	return
}

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum

type Queue []int

func (q *Queue) Push(x int) {
	for q.Size() > 0 && (*q)[q.Size()-1] < x {
		*q = (*q)[:q.Size()-1]
	}

	*q = append(*q, x)
}

func (q *Queue) Peek() int {
	return (*q)[0]
}

func (q *Queue) Pop(x int) {
	if q.Size() > 0 && x == (*q)[0] {
		*q = (*q)[1:]
	}
}

func (q *Queue) Size() int {
	return len(*q)
}

func maxSlidingWindowWithMonotoneStack(nums []int, k int) []int {
	ans := make([]int, 0, len(nums))
	q := Queue{}
	for i := 0; i < k; i++ {
		q.Push(nums[i])
	}
	ans = append(ans, q.Peek())

	for i := k; i < len(nums); i++ {
		q.Pop(nums[i-k])
		q.Push(nums[i])
		ans = append(ans, q.Peek())
	}

	return ans
}

// 239. 滑动窗口最大值
// https://leetcode.cn/problems/sliding-window-maximum
func maxSlidingWindowWithSlice(nums []int, k int) []int {
	// 构造一个非严格单调递减的队列
	q := make([]int, 0, len(nums))
	for i := 0; i < k; i++ {
		q = push(q, nums[i])
	}

	ans := make([]int, 0, len(nums))
	ans = append(ans, q[0])
	for i := k; i < len(nums); i++ {
		q = pop(q, nums[i-k])
		q = push(q, nums[i])
		ans = append(ans, q[0])
	}

	return ans
}

func pop(q []int, num int) []int {
	if q[0] == num {
		q = q[1:]
	}

	return q
}

func push(q []int, num int) []int {
	if len(q) == 0 {
		q = append(q, num)
		return q
	}

	for len(q) > 0 && q[len(q)-1] < num {
		q = q[:len(q)-1]
	}

	q = append(q, num)
	return q
}

// 347. 前 K 个高频元素
// https://leetcode.cn/problems/top-k-frequent-elements
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int, len(nums))
	for _, num := range nums {
		m[num]++
	}

	ans := make([]int, 0, len(m))
	for key := range m {
		ans = append(ans, key)
	}

	sort.Slice(ans, func(i, j int) bool {
		return m[ans[i]] > m[ans[j]]
	})

	return ans[:k]
}
