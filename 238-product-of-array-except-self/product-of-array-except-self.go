func productExceptSelf(nums []int) []int {
    size := len(nums)
    answer := make([]int, size)
    answer[0] = 1
    answer[size-1] = nums[size-1]
    for i:=1; i<size; i++ {
        answer[i] = nums[i-1] * answer[i-1]
        // fmt.Printf("Updated array after %d iteration: %v\n", i, answer)
    }
    rightProduct := 1
    for i:=size-2; i>=0; i-- {
        rightProduct *= nums[i+1]
        answer[i] *= rightProduct
    }
    return answer
}