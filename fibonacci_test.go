package ca_eau

import "testing"

func TestFibonacci(t *testing.T) {
	var v, got, exp int

	v = 0
	got, exp = Fibonacci(v), 0
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 1
	got, exp = Fibonacci(v), 1
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 2
	got, exp = Fibonacci(v), 1
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 3
	got, exp = Fibonacci(v), 2
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 4
	got, exp = Fibonacci(v), 3
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 5
	got, exp = Fibonacci(v), 5
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}

	v = 6
	got, exp = Fibonacci(v), 8
	if got != exp {
		t.Errorf("Fibonacci(%d) = %d, expected = %d", v, got, exp)
	}
}
