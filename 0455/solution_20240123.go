package _455

import "sort"

func findContentChildren(g []int, s []int) int {

	sort.Ints(g)
	sort.Ints(s)

	count := 0
	for i, j := 0, 0; i < len(g); i++ {
		for ; j < len(s); j++ {
			if g[i] <= s[j] {
				count++
				j++
				break
			}
		}
	}
	return count
}
