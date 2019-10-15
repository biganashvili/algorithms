package sorting

//MergeSort realization of merge sort algorithm
func MergeSort(m []int) []int {
	ml := len(m)
	if ml > 1 {
		mid := ml / 2
		l := MergeSort(m[:mid])
		r := MergeSort(m[mid:])
		return merge(l, r)
	}
	return m
}

func merge(l []int, r []int) []int {
	tmp := []int{}
	i := 0
	j := 0
	k := 0
	for i < len(l) && j < len(r) {
		if l[i] < r[j] {
			tmp = append(tmp, l[i])
			i++
		} else {
			tmp = append(tmp, r[j])
			j++
		}
		k++
	}
	for i < len(l) {
		tmp = append(tmp, l[i])
		i++
		k++
	}
	for j < len(r) {
		tmp = append(tmp, r[j])
		j++
		k++
	}
	return tmp
}
