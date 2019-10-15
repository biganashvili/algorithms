package sorting

//QuickSort realization of Quick Sort
func QuickSort(a []int, leftID int, rightID int) {
	if leftID < rightID {
		p := partition(a, leftID, rightID)
		QuickSort(a, leftID, p-1)
		QuickSort(a, p+1, rightID)
	}
}

func partition(a []int, leftID int, rightID int) int {
	index := leftID - 1
	for j := leftID; j < rightID; j++ {
		if a[j] < a[rightID] {
			index++
			a[j], a[index] = a[index], a[j]
		}
	}
	index++
	a[rightID], a[index] = a[index], a[rightID]
	return index
}
