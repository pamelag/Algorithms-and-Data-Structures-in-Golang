package algorithms

func SelectionSort(A []int) []int {
	B := make([]int, 0)
	for i:=0; i <= len(A); i++ {
		minIndex := i
		for j:=i+1; j <=len(A); j++ {
			if A[j] < A[minIndex]{
				minIndex = j
			}
		}
		B = append(B, A[minIndex])
	}
	return B
}

