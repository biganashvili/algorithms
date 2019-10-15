package main

import (
	"algorithms/sorting"
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	slice := randomArray(100000000)

	// fmt.Println(slice)
	start := time.Now()

	sorting.QuickSort(slice, 0, len(slice)-1)
	// fmt.Println(slice)
	fmt.Println("isSorted:", isSorted(slice), " sorted in:", time.Since(start))

	// sorting.BubbleSort(slice)
	// fmt.Println(slice)
	// fmt.Println("isSorted:", isSorted(slice), " sorted in:", time.Since(start))

	// sorted := sorting.MergeSort(slice)
	// fmt.Println(sorted)
	// fmt.Println("isSorted:", isSorted(sorted), " sorted in:", time.Since(start))
}

func randomArray(len int) []int {
	a := make([]int, len)
	for i := 0; i <= len-1; i++ {
		a[i] = rand.Intn(len)
	}
	return a
}

func isSorted(a []int) bool {
	if len(a) < 2 {
		return true
	}
	for i := 0; i < len(a)-1; i++ {
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}
