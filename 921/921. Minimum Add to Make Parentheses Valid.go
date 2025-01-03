package leetcode

import (
	"leetcode/utils"
)

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func minAddToMakeValid(s string) int {
	countParenthesesOpen := 0
	countParenthesesClose := 0
	for _, char := range s {
		if char == '(' {
			countParenthesesOpen++
			continue
		}
		if countParenthesesOpen > 0 {
			countParenthesesOpen--
		}
		countParenthesesClose++

	}
	utils.Log(countParenthesesOpen + countParenthesesClose)
	return countParenthesesOpen + countParenthesesClose
}

func Run(s string) int {

	return minAddToMakeValid(s)

}
