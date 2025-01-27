package leecode

import (
	"strconv"
	"strings"
)

func areNumbersAscending(s string) bool {
	words := strings.Split(s, " ")
	latestNumber := -1
	for i := 0; i < len(words); i++ {

		num, err := strconv.Atoi(words[i])
		if err != nil {
			continue
		}
		if num > latestNumber {
			latestNumber = num
		} else {
			return false
		}
	}
	return true
}

func Run(s string) bool {
	return areNumbersAscending(s)
}
