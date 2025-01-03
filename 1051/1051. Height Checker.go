package leetcode

import "leetcode/utils"

// Complexidade: Tempo O(n^2) no pior caso e O(n) no melhor caso
// Espaço O(n)
func insertionSort(arr []int) {

	for i := 0; i < len(arr); i++ {
		j := i - 1
		key := arr[i]
		for ; j >= 0; j-- {
			if key < arr[j] {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}

		arr[j+1] = key
	}
}

func HeightChecker(heights []int) int {
	heightsSorted := append([]int{}, heights...)
	insertionSort(heightsSorted)
	utils.Log(heightsSorted)
	heightsDiff := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != heightsSorted[i] {
			heightsDiff++
		}
	}
	return heightsDiff
}
