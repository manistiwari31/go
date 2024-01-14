func max(a,b int) int {
    if a>b {return a}
    return b
}
func totalSteps(nums []int) int {
    n := len(nums)
    DP := make([]int, n)
    Stack := make([]int, 1e5)
    top := -1
    ans := 0
    for i := n-1 ; i>=0; i-- {
        for top != -1 && nums[Stack[top]] < nums[i] {
            DP[i] = max(DP[Stack[top]], DP[i]+1)
            top--
        }
        ans = max(ans, DP[i])
        top++
        Stack[top] = i
    }
    return ans
}