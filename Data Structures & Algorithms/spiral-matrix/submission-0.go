func spiralOrder(matrix [][]int) []int {
    m, n := len(matrix), len(matrix[0])
	res := make([]int, 0, m*n)

	// 右、下、左、上四個方向
	dirs := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	d := 0 // 初始方向向右 dir[d] == {0, 1}

	r, c := 0, 0 // 當前座標

	for i := 0; i < m*n; i++ {
		res = append(res, matrix[r][c])
		matrix[r][c] = 200 // 因為原本的值最多就是 100，這代表 200 是走過的位置

		nextR, nextC := r + dirs[d][0], c + dirs[d][1]

		if nextR < 0 || nextR >= m || nextC < 0 || nextC >= n || matrix[nextR][nextC] == 200 { // 代表超過邊界或碰到重複的，要轉向了
			d = (d + 1) % 4
			nextR, nextC = r + dirs[d][0], c + dirs[d][1]
		}

		r, c = nextR, nextC // 繼續走
	}

	return res
}
