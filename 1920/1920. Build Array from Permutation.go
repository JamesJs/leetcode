package leetcode

func buildArray(nums []int) []int {
	prevValue := 0
	for i := 0; i < len(nums); i++ {
		if prevValue == nums[nums[i]] {
			prevValue = nums[i+1]

			continue
		}
		aux := nums[i]
		nums[i] = nums[nums[i]]
		nums[nums[i]] = aux
	}
	return nums
}

func Run(nums []int) []int {
	return buildArray(nums)
}
