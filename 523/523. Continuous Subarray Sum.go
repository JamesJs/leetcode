package ContinuousSubarraySum

func CheckSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	sums := make(map[int]int)
	sum := 0
	nothing := 0
	for i := 0; i < len(nums); i++ {
		//Deal with [0,0] case solution
		if nums[i] == 0 {
			nothing++
			if nothing == 2 {
				return true
			}
			continue
		}
		nothing = 0
		sum = sum + nums[i]

		sums[sum] = i
		_mod := sum % k
		if _mod == 0 {
			return true
		}
		//i-sum[_mod] > 1 is for avoid solutions with only one number
		if _, ok := sums[_mod]; ok && _mod != nums[i] && i-sums[_mod] > 1 {
			return true
		}
	}
	if sum%k == 0 {
		return true
	}

	return false

}
