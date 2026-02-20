func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// 以第一個字串為基準
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		s := strs[i]

		// 找共同長度上限
		minLen := len(prefix)
		if len(s) < minLen {
			minLen = len(s)
		}

		// byte 直接比對，找不同點
		j := 0
		for j < minLen && prefix[j] == s[j] {
			j++
		}

		prefix = prefix[:j]

		// 提前返回：已無公共前綴
		if j == 0 {
			return ""
		}
	}

	return prefix
}