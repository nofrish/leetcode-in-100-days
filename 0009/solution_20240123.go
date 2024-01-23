package _009

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	t := int64(x)

	var y int64
	for x > 0 {
		r := x % 10
		y = y*10 + int64(r)
		x = x / 10
	}

	return y == t
}
