package _009

func isPalindrome____(x int) bool {

	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	y := 0
	for x > y {
		y = y*10 + x%10
		x = x / 10
	}

	if x == y || x == y/10 {
		return true
	}
	return false
}
