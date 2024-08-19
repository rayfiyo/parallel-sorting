package gen

import (
	"math/rand"
	"time"
)

func RandomNumber(n int) []int {
	rand.NewSource(time.Now().UnixNano())

	values := make([]int, n)
	for i := 0; i < n; i++ {
		values[i] = i + 1
	}

	// Fisher-Yates シャッフルアルゴリズムでランダムに並べ替え
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})

	return values
}
