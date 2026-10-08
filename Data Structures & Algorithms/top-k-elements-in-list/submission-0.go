func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)

	for _, num := range nums {
		count[num]++
	}

	buckets := make([][]int, len(nums)+1)
	for num, freq := range count {
		buckets[freq] = append(buckets[freq], num)
	}
	// 從後面開始找，找到 k 個為止
	res := make([]int, 0, k)
	for i := len(buckets) - 1; i > 0; i-- {
		for _, num := range buckets[i] {
			res = append(res, num)
			if len(res) == k {
				return res
			}
		}
	}

	return res

}
