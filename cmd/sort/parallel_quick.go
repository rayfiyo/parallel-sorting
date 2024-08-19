package sort

import (
	"sync"
)

type ParallelQuickSort struct{}

func (pqs ParallelQuickSort) Name() string {
	return "Parallel QuickSort"
}

func (pqs ParallelQuickSort) Sort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	result := make([]int, len(arr))
	copy(result, arr)
	pqs.parallelQuickSort(result, 0, len(result)-1, 4) // 4は初期の深さ、調整可能
	return result
}

func (pqs ParallelQuickSort) parallelQuickSort(arr []int, low, high, depth int) {
	if low < high {
		if depth <= 0 || high-low < 1000 { // 小さな部分配列や深い再帰ではシーケンシャルに
			pqs.sequentialQuickSort(arr, low, high)
			return
		}

		pivot := pqs.partition(arr, low, high)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			pqs.parallelQuickSort(arr, low, pivot-1, depth-1)
		}()

		go func() {
			defer wg.Done()
			pqs.parallelQuickSort(arr, pivot+1, high, depth-1)
		}()

		wg.Wait()
	}
}

func (pqs ParallelQuickSort) sequentialQuickSort(arr []int, low, high int) {
	if low < high {
		pivot := pqs.partition(arr, low, high)
		pqs.sequentialQuickSort(arr, low, pivot-1)
		pqs.sequentialQuickSort(arr, pivot+1, high)
	}
}

func (pqs ParallelQuickSort) partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
