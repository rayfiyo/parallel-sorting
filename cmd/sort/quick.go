package sort

func Quick(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := arr[len(arr)/2]
	left := 0
	right := len(arr) - 1

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	Quick(arr[:left])
	Quick(arr[left:])
}
