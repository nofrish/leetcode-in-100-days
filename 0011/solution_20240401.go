package _011

func maxArea_(height []int) int {

	left, right, result := 0, len(height)-1, 0
	for left < right {
		result = max(result, (right-left)*min(height[left], height[right]))

		// 下面指针变化的思考逻辑是：
		// 1. 让指针指向最两侧，保留所有的选择可能性
		// 2. 指针要开始收缩时，应该让高度较低的指针收缩，因为不论收缩较高的指针得到的下一个高度是多少，都不会比当前圈的面积更大
		// 3. 即高度较低的指针的一次移动，代表着一组可能性的全部消失。而消失的那些可能性，也确实是没有必要考虑的
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return result
}
