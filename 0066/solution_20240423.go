package _066

func plusOne(digits []int) []int {

	result := make([]int, len(digits)+1)
	r := 1
	for i := len(digits) - 1; i >= 0; i-- {
		result[i+1], r = (digits[i]+r)%10, (digits[i]+r)/10
	}
	if r == 0 {
		return result[1:]
	}
	result[0] = r
	return result
}
