func countComponents(n int, edges [][]int) int {
    res := n // 預設每個 components 都是獨立的
	// 每個 components 的 leader 預設都是自己
	leader := make([]int, n)
	for i := 0; i < n; i++ {
		leader[i] = i
	}
	// 找 leader 的函式
	var find func(node int) int
	find = func(node int) int {
		for leader[node] != node {
			// 加速：可以一次跳過兩層查找
			leader[node] = leader[leader[node]]
			node = leader[node]
		}

		return node
	}

	for _, edge := range edges {
		n1, n2 := edge[0], edge[1]
		leader1 := find(n1)
		leader2 := find(n2)
		// 如果不一樣代表有新的 component 被串在一起
		if leader1 != leader2 {
			leader[leader1] = leader2
			res--
		}
	}

	return res
}
