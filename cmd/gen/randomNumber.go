package gen

import (
	"math/rand"
	"time"
)

func RandomNumber(length int) []int {
	rand.NewSource(time.Now().UnixNano())
	arr := make([]int, length)

	for i := range arr {
		// 0から99までのランダムな整数を生成
		arr[i] = rand.Intn(100)
	}

	return arr
}
