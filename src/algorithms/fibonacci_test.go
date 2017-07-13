package algorithms

import "testing"

func TestFibonacci(t *testing.T) {
	err := Fibonacci(20)
	if err != nil {
		t.Error("Error")
	}
}
