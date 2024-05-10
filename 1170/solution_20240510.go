package _170

import "sort"

func numSmallerByFrequency(queries []string, words []string) (result []int) {

	fq := make([]int, 0, len(queries))
	for _, query := range queries { // 显然这里可以抽象为一个函数
		chars := [26]int{}
		for i := 0; i < len(query); i++ {
			chars[query[i]-'a'] += 1
		}
		for i := 0; i < len(chars); i++ {
			if chars[i] != 0 {
				fq = append(fq, chars[i])
				break
			}
		}
	}

	fw := make([]int, 0, len(words))
	for _, word := range words {
		chars := [26]int{}
		for i := 0; i < len(word); i++ {
			chars[word[i]-'a'] += 1
		}
		for i := 0; i < len(chars); i++ {
			if chars[i] != 0 {
				fw = append(fw, chars[i])
				break
			}
		}
	}

	sort.Ints(fw)
	for i := 0; i < len(fq); i++ { // 我应该学习一下 sort.SearchInts 函数的使用
		// 在不是单独考察二分查找的地方，可以用这个函数来快速解决
		lo, hi := 0, len(fw)-1
		for lo < hi {
			mid := lo + (hi-lo)/2
			if fw[mid] <= fq[i] {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		if fw[hi] <= fq[i] {
			result = append(result, 0)
		} else {
			result = append(result, len(fw)-hi)
		}
	}
	return
}
