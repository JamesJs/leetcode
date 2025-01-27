package utils

import (
	"fmt"
)

func LogCharacter(msg rune) {
	fmt.Printf("%c", msg)
}

func Log(msg ...any) {
	fmt.Println(msg...)
	//time.Sleep(5 * time.Second)
}

func CreateArray[K any](n int, m int) [][]K {
	var arr = make([][]K, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]K, m)
	}
	return arr
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
