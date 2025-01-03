package fib

func Fib(n int) int {
	var n0 = 0
	var n1 = 1
	if n == 1 {
		return n1
	}

	response := 0
	for i := 2; i <= n; i++ {
		response = n0 + n1
		n0 = n1
		n1 = response

	}
	return response
}
