package ui

import (
	"context"
	"time"

	"github.com/mum4k/termdash/widgets/barchart"
)

func PlaySort(ctx context.Context, bc *barchart.BarChart, arr []int) error {
	var quickSort func(arr []int, low, high int) error
	quickSort = func(arr []int, low, high int) error {
		if low < high {
			p, err := partition(arr, low, high)
			if err != nil {
				return err
			}

			if err := bc.Values(arr, max); err != nil {
				return err
			}

			// 進行状況を確認できるように適宜遅延を挟む
			time.Sleep(100 * time.Millisecond)

			if err := quickSort(arr, low, p-1); err != nil {
				return err
			}

			if err := quickSort(arr, p+1, high); err != nil {
				return err
			}
		}
		return nil
	}

	return quickSort(arr, 0, len(arr)-1)
}

func partition(arr []int, low, high int) (int, error) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i, nil
}
