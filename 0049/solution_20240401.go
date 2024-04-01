package _049

func groupAnagrams__(strs []string) [][]string {

	anagrams := make(map[[26]byte][]string)
	for _, str := range strs {
		key := [26]byte{}
		for i := 0; i < len(str); i++ {
			key[str[i]-'a'] += 1
		}
		anagrams[key] = append(anagrams[key], str)
	}

	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}
	return result
}
