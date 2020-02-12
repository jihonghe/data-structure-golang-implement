package utils

import (
	"encoding/json"
	"io/ioutil"
)

type NumsData struct {
	ints []int
}

func WriteIntsToFile(nums []int, filePath string) {
	numsBytes, _ := json.Marshal(nums)
	ioutil.WriteFile(filePath, numsBytes, 0644)
}

func ReadIntsFromFile(filePath string) []int {
	var ints []int
	intsBytes, _ := ioutil.ReadFile(filePath)

	json.Unmarshal(intsBytes, &ints)
	println("from file length:", len(ints))

	return ints
}

func IntsToints64(ints []int) []int64 {
	intsLength := len(ints)
	ints64 := make([]int64, intsLength, intsLength)
	for index, value := range ints {
		ints64[index] = int64(value)
	}

	return ints64
}

func ReadInt64sFromFile(filePath string) []int64 {
	var ints []int64
	intsBytes, _ := ioutil.ReadFile(filePath)

	json.Unmarshal(intsBytes, &ints)

	return ints
}
