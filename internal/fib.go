package internal

func fib(n int) int {
	f1, f2 := 0, 1
	if n == 0 {
		return f1
	}
	if n == 1 {
		return f2
	}
	var ret int
	for i := 2; i <= n; i++ {
		ret = (f1 + f2) % (1e9 + 7)
		f1 = f2
		f2 = ret
	}
	return ret
}
