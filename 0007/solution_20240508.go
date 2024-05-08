package _007

func reverse(x int) int {

	var res int32
	for x != 0 {
		last := res
		res = res*10 + int32(x)%10
		x = x / 10
		if res/10 != last {
			return 0
		}
	}
	return int(res)
}
