func threeSum(nums []int) [][]int {
	sort.Ints(nums)
    res := [][]int{}
    n := len(nums)

    for i := 0; i < n-2; i++ {
        // nums 已經排序，若當前基準大於 0，後面三數之和必大於 0
        if nums[i] > 0 {
            break
        }
        // 避免第一個數字重複
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }

        l, r := i+1, n-1
        for l < r {
            sum := nums[i] + nums[l] + nums[r]
            if sum == 0 {
                res = append(res, []int{nums[i], nums[l], nums[r]})
                l++
                r--
                // 避開重複的第二個與第三個數字
                for l < r && nums[l] == nums[l-1] {
                    l++
                }
                for l < r && nums[r] == nums[r+1] {
                    r--
                }
            } else if sum < 0 {
                l++
            } else {
                r--
            }
        }
    }

    return res
}
