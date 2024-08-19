package sort

type Sorter interface {
	Sort([]int) []int
	Name() string
}
