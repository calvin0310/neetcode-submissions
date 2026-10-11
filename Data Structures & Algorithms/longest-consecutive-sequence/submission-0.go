func longestConsecutive(nums []int) int {
	maxStreak := 0
	// struct{} 不佔用記憶體
	numSet := make(map[int]struct{}, len(nums))
	// 先記錄每個數字
	for _, num := range nums {
		numSet[num] = struct{}{}
	}

	// 尋找起點，然後開始計算最長連續數字
	for num := range numSet {
		if _, existed := numSet[num - 1]; existed {
			// 不是起點
			continue
		}

		streak := 0
		for i := num;; i++{
			if _, existed := numSet[i]; existed {
				streak += 1
			} else {
				break
			}
		}

		if streak > maxStreak {
			maxStreak = streak
		} 
	}

	return maxStreak
}
