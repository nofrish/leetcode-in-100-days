package _009

func isPalindrome___(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	y := 0
	for y < x {
		y = y*10 + x%10
		x = x / 10
	}
	return x == y || x == y/10
}
