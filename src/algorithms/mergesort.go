package algorithms

import "errors"

func MergeSort(A []string) ([]string, error) {
	if len(A) == 1 {
		return errors.New("Empty set")
	}
	size := len(A)
	if (size) == 1 {
		return A
	}
	A1 := make([]string, 0)
	mid := size / 2

	a := sort(A[0:mid])
	a1 := sort(A[mid:size])
	A1 = merge(a, a1)
	return A1, nil
}

func sort(param []string) []string {
	B := make([]int, 0)
	for i:=0; i <= len(param); i++ {
		minIndex := i
		for j:=i+1; j <=len(param); j++ {
			if param[j] < param[minIndex]{
				minIndex = j
			}
		}
		B = append(B, param[minIndex])
	}
	return B
}

func merge(a []string, a1 []string) []string {
	D := make([]string, len(a)+len(a1))
	i := 0
	// This loop contributes while a and a1 both non empty
	for (len(a) > 0 && len(a1) > 0) {
		m := a[i]
		n := a1[i]
		if m < n {
			D = append(D, a[m:])
		} else {
			D = append(D, a[n:])
		}
	}
	if len(a) > 0 {
		D = append(D, a[:])
	} else {
		D = append(D, a1[:])
	}
	return D
}