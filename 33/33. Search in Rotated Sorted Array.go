package SearchRotate

//disable codeium
import (
	utils "leetcode/utils"
)

func isIntOdd(num int) bool {
	if num%2 == 0 {
		return false
	}
	return true
}

func binarySearch(arr []int, start int, end int, search int) int {
	if start >= end || start >= len(arr) || end < 0 {
		return -1
	}

	if arr[start] == search {
		return start
	}
	if arr[end] == search {
		return end
	}

	size := end - start + 1
	middle := ((size) / 2) + start

	utils.Log(size, middle, start, end)
	if middle >= len(arr) {
		return -1
	}
	if arr[middle] == search {
		return middle
	}

	//linear para esquerda
	if arr[start] < arr[middle] {
		if search > arr[start] && search < arr[middle] {
			return binarySearch(arr, start, middle-1, search)
		}
		return binarySearch(arr, middle+1, end-1, search)
	}
	if search > middle {
		return binarySearch(arr, middle+1, end-1, search)
	}
	return binarySearch(arr, start+1, middle-1, search)
}

func Search(nums []int, target int) int {
	size := len(nums)
	if size == 1 && nums[0] != target {
		return -1
	}
	if size == 2 && nums[0] != target && nums[1] != target {
		return -1
	}
	utils.Log(binarySearch(nums, 0, len(nums)-1, target))
	return 0
}
