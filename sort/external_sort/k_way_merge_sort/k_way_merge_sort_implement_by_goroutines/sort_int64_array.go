package k_way_merge_sort_implement_by_goroutines

import "fmt"

// 实现了sort.interface接口，以便调用sort.Sort(data interface)
type Int64Array []int64

func NewInt64Array (ints... int) Int64Array {
	int64Array := make(Int64Array, len(ints))
	for index, intValue := range ints {
		int64Array[index] = int64(intValue)
	}

	return int64Array
}

func TraverseInt64Array(array []int64) {
	for _, value := range array {
		fmt.Printf("%d, ", value)
	}
	println()
}

func (int64Array Int64Array) Len() int {
	return len(int64Array)
}

func (int64Array Int64Array) Less(i, j int) bool {
	return int64Array[i] < int64Array[j]
}

func (int64Array Int64Array) Swap(i, j int) {
	int64Array[i], int64Array[j] = int64Array[j], int64Array[i]
}
