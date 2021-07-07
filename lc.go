/*238. 除自身以外数组的乘积*/
func productExceptSelf(nums []int) []int {
    l := len(nums)
    res := make([]int,l)
    for i,_ := range res {
        res[i] = 1
    }
    left,right := 1,1
    for i,_ := range nums {
        res[i] *= left
        res[l-i-1] *= right
        left *= nums[i]
        right *= nums[l-i-1]
    }
    return res
}
