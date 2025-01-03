package leetcode

import "leetcode/utils"

func waysToSplitArray(nums []int) int {
	i := 0
	_len := len(nums) - 1
	totalSum := 0
	for i <= _len {
		totalSum += nums[i]
		i++
	}
	i = _len - 1
	localSum := nums[_len]
	total := 0
	for i >= 0 {
		if totalSum-localSum >= localSum {
			total++
		}
		localSum += nums[i]
		i--
	}
	return total
}

func Run(nums []int) int {
	return waysToSplitArray(nums)
}
