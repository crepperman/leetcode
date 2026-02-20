func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	n := len(s)
	result := make([]byte, n) // 直接分配完整大小，不用 append
	cycle := 2 * (numRows - 1)
	base := unsafe.Pointer(unsafe.StringData(s))
	pos := 0

	for row := 0; row < numRows; row++ {
		for j := 0; j+row < n; j += cycle {
			result[pos] = *(*byte)(unsafe.Pointer(uintptr(base) + uintptr(j+row)))
			pos++

			if row != 0 && row != numRows-1 {
				mid := j + cycle - row
				if mid < n {
					result[pos] = *(*byte)(unsafe.Pointer(uintptr(base) + uintptr(mid)))
					pos++
				}
			}
		}
	}

	return unsafe.String(&result[0], n)
}