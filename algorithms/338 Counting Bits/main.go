func countBits(n int) []int {
	ans := make([]int, n+1)

	for i := 1; i <= n; i++ {
		// i>>1 = i/2，右移一位去掉最低位
		// i&1  = 最低位是否為1
		ans[i] = ans[i>>1] + (i & 1)
	}

	return ans
}