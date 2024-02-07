package _135

func candy(ratings []int) int {

	counts := make([]int, len(ratings))
	for i := 0; i < len(ratings); i++ {
		counts[i] = 1
	}

	for {
		changed := false
		for i := 0; i < len(ratings); i++ {
			if i > 0 && ratings[i] > ratings[i-1] && counts[i] <= counts[i-1] {
				counts[i] = counts[i-1] + 1
				changed = true
			}
			if i < len(ratings)-1 && ratings[i] > ratings[i+1] && counts[i] <= counts[i+1] {
				counts[i] = counts[i+1] + 1
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	result := 0
	for i := 0; i < len(counts); i++ {
		result += counts[i]
	}
	return result
}
