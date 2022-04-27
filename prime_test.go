package water

import "testing"

func TestNextPrime(t *testing.T) {
	var v, exp int
	v, exp = 1, 2
	if got := NextPrime(v); got != exp {
		t.Errorf("NextPrime(%d) = %d; expected = %d", v, exp, got)
	}

	v, exp = 2, 3
	if got := NextPrime(v); got != exp {
		t.Errorf("NextPrime(%d) = %d; expected = %d", v, exp, got)
	}

	v, exp = 3, 5
	if got := NextPrime(v); got != exp {
		t.Errorf("NextPrime(%d) = %d; expected = %d", v, exp, got)
	}

	v, exp = 14, 17
	if got := NextPrime(v); got != exp {
		t.Errorf("NextPrime(%d) = %d; expected = %d", v, exp, got)
	}
}
