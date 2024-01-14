func max(a,b int) int {
    if a>b {return a}
    return b
}
func totalSteps(nums []int) int {
    steps := 0
    ans := 0
	stack := [][2]int{}
    for i := len(nums)-1 ; i >= 0; i-- {
        steps = 0
        for len(stack) > 0 {
            item := stack[len(stack)-1]
            if nums[i] > item[0] {
                steps = max(steps+1, item[1])          
                stack = stack[:len(stack)-1]
            } else {
                break
            }
        }
        
        ans = max(ans, steps)
        stack = append(stack, [2]int{nums[i], steps})
        
    }
    return ans
}