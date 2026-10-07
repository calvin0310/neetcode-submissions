func reverseBits(n int) int {
	res := 0

	for i := 0; i < 32; i++ {
		bit := n & 1 //  最右邊的一位

		// 往左推一位 加上 n 最右邊的一位
		res = (res << 1) | bit
		// 往右推一位
		n >>= 1
	}

	return res
}
