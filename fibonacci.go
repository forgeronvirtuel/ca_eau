package water

func Fibonacci(n int) int {

	f_2 := 0
	f_1 := 1

	if n < 0 {
		return -1
	}
	if n == 0 {
		return f_2
	}

	if n == 1 {
		return f_1
	}

	for i := 2; i <= n; i++ {
		f_1, f_2 = f_1+f_2, f_1
	}

	return f_1
}
