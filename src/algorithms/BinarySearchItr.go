package algorithms

import "errors"

func BinarySearchItr(A []int, high int, low int, key int) (int, error) {
	mid := 0
	for low <= high {
		mid = low + (high-low)/2
		if key == A[mid] {
			return mid, nil
		} else if key > mid {
			low = low + mid
		} else {
			high = mid - 1
		}
	}
	return low - 1, errors.New("key not found")
}