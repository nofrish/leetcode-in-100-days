package _011

func maxArea(height []int) int {

	low, high, max := 0, len(height)-1, 0
	for low < high {
		area := (high - low) * min(height[low], height[high])
		if max < area {
			max = area
		}

		if height[low] < height[high] {
			low++
		} else {
			high--
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
