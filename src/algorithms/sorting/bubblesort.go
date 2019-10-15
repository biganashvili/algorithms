package sorting

import "fmt"

//BubbleSort realization of bubble sort
func BubbleSort(a []int) {
	isSorted := false
	k := 0
	for !isSorted {
		isSorted = true
		l := len(a)
		fmt.Println(k)
		for j := 0; j < l-k-1; j++ {
			if a[j] > a[j+1] {
				isSorted = false
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		k++
	}
}
