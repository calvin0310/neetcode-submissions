func maxArea(heights []int) int {
    l, r := 0, len(heights)-1
    maxWater := 0

    for l < r {
        // 計算當前高度與寬度
        h := min(heights[l], heights[r])
        width := r - l
        currentWater := h * width

        if currentWater > maxWater {
            maxWater = currentWater
        }

        // 移動較矮的一側
        if heights[l] < heights[r] {
            l++
        } else {
            r--
        }
    }

    return maxWater
}