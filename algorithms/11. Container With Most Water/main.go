func maxArea(height []int) int {
	left, right := 0, len(height)-1
	max := 0

	for left < right {
		hl := height[left]
		hr := height[right]

		// 無分支取 min
		h := hl
		if hr < h {
			h = hr
		}

		area := h * (right - left)
		if area > max {
			max = area
		}

		// 同時移動，減少分支判斷
		if hl < hr {
			left++
			// 跳過比 hl 更矮的，無意義的移動
			for left < right && height[left] <= hl {
				left++
			}
		} else {
			right--
			// 跳過比 hr 更矮的，無意義的移動
			for left < right && height[right] <= hr {
				right--
			}
		}
	}

	return max
}