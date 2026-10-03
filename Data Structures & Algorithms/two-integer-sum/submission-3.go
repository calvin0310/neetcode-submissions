func twoSum(nums []int, target int) []int {
    record := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if j, ok := record[diff]; ok {
			return []int{j, i}
		}

		record[num] = i
	}

	return nil
}
