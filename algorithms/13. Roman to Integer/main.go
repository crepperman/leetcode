func romanToInt(s string) int {
	// 查表：直接用 byte 當 index，省去 map lookup
	val := [128]int{}
	val['I'] = 1
	val['V'] = 5
	val['X'] = 10
	val['L'] = 50
	val['C'] = 100
	val['D'] = 500
	val['M'] = 1000

	result := 0
	n := len(s)

	for i := 0; i < n-1; i++ {
		cur := val[s[i]]
		next := val[s[i+1]]

		// 當前比下一個小->減法
		if cur < next {
			result -= cur
		} else {
			result += cur
		}
	}

	// 最後一個字符永遠加
	result += val[s[n-1]]

	return result
}