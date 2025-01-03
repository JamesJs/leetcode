package trifib

func help(n int, cache map[int]int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	if _, has := cache[n]; has {
		return cache[n]
	}
	cache[n] = help(n-1, cache) + help(n-2, cache) + help(n-3, cache)
	return cache[n]
}

func tribonacci(n int) int {
	var cache map[int]int
	cache = make(map[int]int)
	return help(n, cache)

}
