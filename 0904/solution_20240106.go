package _904

func totalFruit(fruits []int) int {

	slow, fast, bs, max := 0, -1, make(map[int]int, 0), 0
	for fast++; fast < len(fruits); fast++ {
		fruit := fruits[fast]
		if len(bs) < 2 || bs[fruit] > 0 {
			bs[fruit] += 1
			if max < fast-slow+1 {
				max = fast - slow + 1
			}
			continue
		}

		for ; len(bs) == 2; slow++ {
			bs[fruits[slow]] -= 1
			if bs[fruits[slow]] == 0 {
				delete(bs, fruits[slow])
			}
		}

		bs[fruit] += 1
	}

	return max
}
