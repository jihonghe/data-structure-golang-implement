package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "a你好"
	c := "你"

	println(sliceOperation([]string{"a", "b", "c"}))
	fmt.Printf("%s\n", basename("b/hello.go"))
	fmt.Printf("%s\n", basename1("b/hello.go"))
	fmt.Printf("%s\n", string(s[0]))
	fmt.Printf("%s\n", string(s[1:4]))
	fmt.Printf("length: %d\n", len(c))

	for index, value := range s {
		fmt.Printf("index: %d, value: %s\n", index, string(value))
	}

}

func sliceOperation(s []string) bool {
	out := s[:0]
	if &out == &s {
		return true
	}
	return false
}

func basename1(s string) string {
	var targetStr string

	lastSlashIndex := strings.LastIndex(s, "/")
	targetStr = s[lastSlashIndex + 1:]
	if dot := strings.LastIndex(targetStr, "."); dot >= 0 {
		targetStr = targetStr[:dot]
	}

	return targetStr
}

func basename(s string) string {
	var targetStr string
	for index, value := range s {
		if value == '/' {
			targetStr = s[index + 1:]
			break
		}
	}
	for i := len(targetStr) - 1; i >= 0; i-- {
		if targetStr[i] == '.' {
			targetStr = targetStr[:i]
			break
		}
	}

	return targetStr
}
