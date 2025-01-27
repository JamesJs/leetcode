package leetcode

import (
	"sort"
	"strings"
)

//-------------Two Hashes-----------------------

func createHashmapFromString(s string, h map[byte]int) {
	for i := 0; i < len(s); i++ {
		if _, ok := h[s[i]]; ok {
			h[s[i]] += 1
			continue
		}
		h[s[i]] = 1
	}

}

func isAnagramTwoHashes(s string, t string) bool {
	var h1 map[byte]int = make(map[byte]int)
	var h2 map[byte]int = make(map[byte]int)
	if len(s) != len(t) {
		return false
	}
	createHashmapFromString(s, h1)
	createHashmapFromString(t, h2)
	for key := range h1 {
		if _, ok := h2[key]; !ok {
			return false
		}
		if h2[key] != h1[key] {
			return false
		}
	}
	return true
}

//---------------SORT---------------------

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i int, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	s = sortString(s)
	t = sortString(t)
	if strings.Compare(s, t) != 0 {
		return false
	}
	return true
}

//---------------One hash---------------------

func isAnagramOneHash(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	h := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		h[s[i]]++
		h[t[i]]--
	}
	for _, value := range h {
		if value != 0 {
			return false
		}
	}
	return true
}

func Run(s string, t string) bool {
	return isAnagramSort(s, t)
}
