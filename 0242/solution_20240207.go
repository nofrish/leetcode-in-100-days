package _242

func isAnagram(s string, t string) bool {

	charsOfS := [26]int{}
	charsOfT := [26]int{}

	for i := 0; i < len(s); i++ {
		charsOfS[int(s[i]-'a')] += 1
	}
	for i := 0; i < len(t); i++ {
		charsOfT[int(t[i]-'a')] += 1
	}

	return charsOfS == charsOfT
}
