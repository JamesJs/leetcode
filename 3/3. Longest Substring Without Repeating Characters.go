package leetcode

import (
	"leetcode/utils"
)

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var m map[rune]int = make(map[rune]int)
	globalMax := 0
	actualMax := 0
	for pos, character := range s {

		posActualCharacter, isIn := m[character]
		if !isIn {
			actualMax++
			m[character] = pos
			continue
		}
		if actualMax > globalMax {
			globalMax = actualMax
		}
		for character, posMap := range m {
			if posMap < posActualCharacter {
				delete(m, character)
				actualMax--
			}

		}
		m[character] = pos

	}

	utils.Log(actualMax, globalMax)
	if actualMax > globalMax {
		return actualMax
	}
	return globalMax
}
func Run(s string) int {
	return lengthOfLongestSubstring(s)

}
