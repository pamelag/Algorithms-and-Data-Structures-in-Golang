package algorithms

import (
	"errors"
	"fmt"
)

func Fibonacci(size int) error {
	if size == 0 {
		return errors.New("Input is empty")
	}
	numbers := make([]int, size)
	numbers[0] = 0
	numbers[1] = 1
	for i := 2; i < size; i++ {
		numbers[i] = numbers[i-1] + numbers[i-2]
	}
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	return nil
}
