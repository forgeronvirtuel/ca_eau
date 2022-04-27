package water

func isPrime(n int) bool {
	for div := 2; div < n; div++ {
		if n%div == 0 {
			return false
		}
	}
	return true
}

// NextPrime finds the next prime number after the given number.
func NextPrime(n int) int {
	var prime int
	for nxt := n + 1; prime == 0; nxt++ {
		if isPrime(nxt) {
			prime = nxt
		}
	}
	return prime
}
