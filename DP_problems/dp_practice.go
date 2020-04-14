package DP_problems

// ------------------------------------------------Practice1 start------------------------------------------------

/*
    用给定的几个数，凑出指定的值，要求使用的数尽可能少。
    给定：1， 5， 11
    凑出的值：15

    使用动态规划的思想，问题可以转化为求数学表达式：f(n) = min{f(n-1), f(n-5), f(n-11} + 1，f(n)的值即为使用给定的数的次数
 */

func MinCount(nums []int, v int) int {
    tmpInts := make([]int, v + 1, v + 1)
    tmpInts[0] = 0
    for index := 1; index <= v; index++ {
        minCost := v + 1
        for _, num := range nums {
            // 剪枝
            if index - num >= 0 {
                minCost = min(minCost, tmpInts[index - num] + 1)
            }
        }
        tmpInts[index] = minCost
    }

    return tmpInts[v]
}

func min(num1, num2 int) int {
    if num1 <= num2 {
        return num1
    }
    return num2
}
// ------------------------------------------------Practice1 end------------------------------------------------

// ------------------------------------------------Practice2 start------------------------------------------------
/*
    最长上升子序列长度问题：
        最长上升子序列定义：子序列的元素在原序列中不一定相邻，满足单调递增即可
        解法：一种DP是O(n^2)，一种DP是O(NlgN)
 */

func MaxSubIncSeqByDP1(seq []int) int {
    seqLen := len(seq)
    dp := make([]int, seqLen, seqLen)
    for i := 0; i < seqLen; i++ {
        dp[i]++
        for j := 0; j < i; j++ {
            if seq[j] < seq[i] && dp[i] < dp[j] + 1 {
                dp[i] = dp[j] + 1
            }
        }
    }
    maxLen := 0
    for i := 0; i < seqLen; i++ {
        if maxLen < dp[i] {
            maxLen = dp[i]
        }
    }

    return maxLen
}

