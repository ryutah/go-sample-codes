package main

func fib1(n int64) int64 {
	var fibPair func(int64) (int64, int64)
	fibPair = func(k int64) (int64, int64) {
		if k == 1 {
			return 1, 0
		}
		i, j := fibPair(k - 1)
		return i + j, i
	}
	i, _ := fibPair(n)
	return i
}

func fib2(n int64) int64 {
	if n <= 2 {
		return 1
	}
	return fib2(n-1) + fib2(n-2)
}
