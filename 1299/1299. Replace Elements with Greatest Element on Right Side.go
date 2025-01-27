package leetcode

import (
	"fmt"
	"sort"
)

func replaceElements(arr []int) []int {
	size := len(arr)
	result := make([]int, size)
	result[size-1] = -1
	max := arr[size-1]
	for i := size - 2; i >= 0; i-- {
		result[i] = max
		if arr[i] > max {
			max = arr[i]
		}
	}
	return result
}

func replaceElementsSort(arr []int) []int {
	size := len(arr)
	h := make([][]int, size)
	for i := 0; i < size; i++ {
		h[i] = []int{i, arr[i]}
	}
	sort.Slice(h, func(i int, j int) bool {
		return h[i][1] > h[j][1]
	})
	fmt.Print(h)
	result := make([]int, size)
	i := 0
	j := 0
	for i < size {
		if i == size-1 {
			result[i] = -1
			break
		}
		if i >= h[j][0] {
			if j >= size {
				break
			}
			j++
			continue
		}
		result[i] = h[j][1]
		i++

	}
	return result

}

func Run(arr []int) []int {
	return replaceElementsSort(arr)
}
