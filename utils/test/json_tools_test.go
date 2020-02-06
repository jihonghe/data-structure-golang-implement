package test

import (
	"data-structure-golang-implement/utils"
	"fmt"
	"testing"
)

func TestWriteIntsToJsonFile(t *testing.T) {
	utils.WriteIntsToFile([]int{1, 2, 3, 4}, "nums0")
}

func TestReadIntsFromFile(t *testing.T) {
	ints := utils.ReadIntsFromFile("/home/jihonghe/gitRepos/data-structure-golang-implement/sort/test/external_sort/k_way_merge_sort/nums")
	fmt.Printf("type: %T, length: %d\n", ints, len(ints))
}
