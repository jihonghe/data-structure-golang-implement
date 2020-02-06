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

	return ints
}
