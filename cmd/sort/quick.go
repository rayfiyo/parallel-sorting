package sort

type QuickSort struct{}

func (qs QuickSort) Name() string {
	return "QuickSort"
}

func (qs QuickSort) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[len(arr)/2]
	left := make([]int, 0)
	right := make([]int, 0)
	equal := make([]int, 0)

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}

	left = qs.Sort(left)
	right = qs.Sort(right)

	return append(append(left, equal...), right...)
}
