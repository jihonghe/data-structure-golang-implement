package DP_problems

import (
    "fmt"
    "testing"
)

func TestMinCount(t *testing.T) {
    nums := []int{2, 5, 7}
    v := 27
    fmt.Printf("cost nums count: %d\n", MinCount(nums, v))
}

func TestMaxSubIncSeq(t *testing.T) {
    nums := []int{9, 0, -1, -2, 3, -9, -8, -5, -4, 1, 1, 5, 4, 2, 6, 8, 10, -3, 7}
    maxSubIncSeq := MaxSubIncSeqByDP1(nums)
    println()
    fmt.Printf("%d\n", maxSubIncSeq)
}
