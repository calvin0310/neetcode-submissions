func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := [][]int{}
    
    for i := 0; i < len(nums) - 2; i++ {
        if nums[i] > 0 {
            break // 後面的數都是遞增，不可能三者加在一起為 0
        }

        if i > 0 && nums[i] == nums[i-1] {
            continue // 跳過重複數
        }


        left, right := i+1, len(nums)-1

        for left < right {
            sum := nums[i] + nums[left] + nums[right]

            if sum == 0 {
                res = append(res, []int{nums[i], nums[left], nums[right]})

                for left < right && nums[left] == nums[left+1] {
                    left++ // 跳過重復項
                }

                for left < right && nums[right] == nums[right-1] {
                    right-- // 跳過重復項
                }

                left++
                right--
            } else if sum < 0 {
                left++
            } else {
                right--
            }
        }
    }

    return res
}
