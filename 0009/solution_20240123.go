package _009

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	t := int64(x)

	var y int64
	for x > 0 {
		y = y*10 + int64(x%10)
		x = x / 10
	}

	return y == t
}

func isPalindrome_(x int) bool {

	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	t := 0
	for x > t {
		t = t*10 + x%10
		x = x / 10
	}
	return x == t || x == t/10
}
