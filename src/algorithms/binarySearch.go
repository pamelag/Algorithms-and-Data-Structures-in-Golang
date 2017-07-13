package algorithms

import "errors"

func BinarySearch(A []int, low int, high int, key int) (int,error) {
	// if high is less than row
	if high < low {
		return errors.New("High is greater than low")
	}
	mid := low + (low + high)/2
	if key == A[mid] {
		return mid, nil
	} else if key < A[mid] {
		BinarySearch(A, low, low+mid, key)
	} else {
		BinarySearch(A, mid+1, high, key)
	}
	return low-1, nil
}
