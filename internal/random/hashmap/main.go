package hashmap

import (
	"fmt"
	"sort"
)

// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram
func isAnagram(s string, t string) bool {
	// 使用 hashmap 判断两者字母数量是否一致
	if len(s) != len(t) {
		return false
	}

	cntMap := make(map[byte]int, len(s))
	for _, b := range []byte(s) {
		if cnt, ok := cntMap[b]; ok {
			cntMap[b] = cnt + 1
		} else {
			cntMap[b] = 1
		}
	}

	for _, b := range []byte(t) {
		if cnt, ok := cntMap[b]; ok {
			cntMap[b] = cnt - 1
		} else {
			return false
		}
	}

	for _, cnt := range cntMap {
		if cnt != 0 {
			return false
		}
	}

	return true
}

// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram
func isAnagramWithOnceLoop(s string, t string) bool {
	// 使用一次循环，记录并完成统计
	if len(s) != len(t) {
		return false
	}

	cntMap := make(map[byte]int, len(s))
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i, j = i+1, j-1 {
		sCnt, _ := cntMap[s[i]]
		cntMap[s[i]] = sCnt + 1

		tCnt, _ := cntMap[t[j]]
		cntMap[t[j]] = tCnt - 1
	}

	for _, cnt := range cntMap {
		if cnt != 0 {
			return false
		}
	}

	return true
}

// 242. 有效的字母异位词
// https://leetcode.cn/problems/valid-anagram
func isAnagramWithArr(s string, t string) bool {
	// 因为只包括小写字母，可以使用数组来实现
	if len(s) != len(t) {
		return false
	}

	charCnt := make([]int, 26)
	for i := range s {
		if s[i] == t[i] {
			continue
		}

		sIdx := s[i] - 'a'
		charCnt[sIdx]++

		tIdx := t[i] - 'a'
		charCnt[tIdx]--
	}

	for _, cnt := range charCnt {
		if cnt != 0 {
			return false
		}
	}

	return true
}

// 349. 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays
func intersection(nums1 []int, nums2 []int) []int {
	numsMap := make(map[int]struct{}, len(nums1))
	for _, num := range nums1 {
		numsMap[num] = struct{}{}
	}

	resMap := make(map[int]struct{}, len(nums2))
	for _, num := range nums2 {
		_, ok := numsMap[num]
		if ok {
			resMap[num] = struct{}{}
		}
	}

	res := make([]int, 0, len(resMap))
	for num := range resMap {
		res = append(res, num)
	}

	return res
}

// 349. 两个数组的交集
// https://leetcode.cn/problems/intersection-of-two-arrays
func intersectionWithArr(nums1 []int, nums2 []int) []int {
	// 因为数组中数字范围是 [0, 1000]
	nums := make([]int, 1001)
	for _, num := range nums1 {
		nums[num]++
	}

	resMap := make(map[int]struct{}, len(nums2))
	for _, num := range nums2 {
		if nums[num] > 0 {
			resMap[num] = struct{}{}
		}
	}

	res := make([]int, 0, len(resMap))
	for num := range resMap {
		res = append(res, num)
	}

	return res
}

// 202. 快乐数
// https://leetcode.cn/problems/happy-number
func isHappy(n int) bool {
	// 使用 bool 数组方便判断
	set := make(map[int]bool)
	for n != 1 && !set[n] {
		// 将这个 n 保存下来
		set[n] = true
		// 计算下一个 n
		n = calcNextNum(n)
	}

	return n == 1
}

func calcNextNum(num int) int {
	sum := 0
	for num > 0 {
		sum += (num % 10) * (num % 10)
		num = num / 10
	}

	return sum
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum
func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int, len(nums))
	for i, num := range nums {
		if val, ok := numsMap[target-num]; ok {
			return []int{val, i}
		}

		numsMap[num] = i
	}

	return []int{}
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum
func twoSumWithTraverse(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return []int{}
}

// 1. 两数之和
// https://leetcode.cn/problems/two-sum
func twoSumWithSort(nums []int, target int) []int {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)

	sort.Ints(nums)

	left, right := 0, len(nums)-1
	for left < right {
		sum := nums[left] + nums[right]
		if sum < target {
			left++
		} else if sum > target {
			right--
		} else {
			break
		}
	}

	for i := 0; i < len(nums); i++ {
		if copiedNums[i] == nums[left] {
			left = i
			break
		}
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if copiedNums[i] == nums[right] {
			right = i
			break
		}
	}

	return []int{left, right}
}

// 454. 四数相加 II
// https://leetcode.cn/problems/4sum-ii
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 因为最后返回的是数量，所以可以分为两组计算
	length := len(nums1)
	sumMap1 := make(map[int]int, length)
	sumMap2 := make(map[int]int, length)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			sum1 := nums1[i] + nums2[j]
			sumMap1[sum1]++

			sum2 := nums3[i] + nums4[j]
			sumMap2[sum2]++
		}
	}

	res := 0
	for sum1, cnt1 := range sumMap1 {
		if cnt2, ok := sumMap2[0-sum1]; ok {
			res += cnt2 * cnt1
		}
	}

	return res
}

// 383. 赎金信
// https://leetcode.cn/problems/ransom-note
func canConstruct(ransomNote string, magazine string) bool {
	charCnt := make(map[byte]int, len(magazine))
	for _, char := range []byte(magazine) {
		charCnt[char]++
	}

	for _, char := range []byte(ransomNote) {
		if cnt, ok := charCnt[char]; ok {
			cnt--
			if cnt < 0 {
				return false
			}
			charCnt[char] = cnt
			continue
		}

		return false
	}

	return true
}

// 383. 赎金信
// https://leetcode.cn/problems/ransom-note
func canConstructWithArr(ransomNote string, magazine string) bool {
	// 因为全部为小写字母构成，可以使用数组
	lettersCnt := make([]int, 26)
	for _, char := range magazine {
		idx := char - 'a'
		lettersCnt[idx]++
	}

	for _, char := range ransomNote {
		idx := char - 'a'
		if lettersCnt[idx] > 0 {
			lettersCnt[idx] = lettersCnt[idx] - 1
		} else {
			return false
		}
	}

	return true
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums))
	resMap := make(map[string]bool, len(nums))
	for i, num := range nums {
		if num > 0 {
			continue
		}

		for j, k := i+1, len(nums)-1; j < len(nums) && k > j; {
			sum := num + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				key := fmt.Sprintf("%d-%d-%d", num, nums[j], nums[k])
				if !resMap[key] {
					res = append(res, []int{num, nums[j], nums[k]})
					resMap[key] = true
				}
				k--
				j++
			}
		}
	}

	return res
}

// 15. 三数之和
// https://leetcode.cn/problems/3sum
func threeSumWithDoublePtr(nums []int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums))
	for i, num := range nums {
		if num > 0 {
			continue
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j, k := i+1, len(nums)-1; j < len(nums) && k > j; {
			sum := num + nums[j] + nums[k]
			if sum > 0 {
				k--
			} else if sum < 0 {
				j++
			} else {
				res = append(res, []int{num, nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				k--
				j++
			}
		}
	}

	return res
}

// 18. 四数之和
// https://leetcode.cn/problems/4sum
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)

	res := make([][]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		if target > 0 && nums[i] > target {
			// 剪枝
			break
		}

		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			// 剪枝
			if nums[i]+nums[j] > target && target > 0 {
				break
			}

			// 去重
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, len(nums)-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum > target {
					right--
				} else if sum < target {
					left++
				} else {
					res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
					// 去重
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					// 去重
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				}
			}
		}
	}

	return res
}
