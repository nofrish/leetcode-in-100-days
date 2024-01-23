package _005

func longestPalindrome(s string) string {

	result := s[0:1]
	for i := 1; i < len(s); i++ { // case "aba"
		for j := 1; i-j >= 0 && i+j < len(s); j++ {
			if s[i-j] == s[i+j] {
				if len(result) < 2*j+1 {
					result = s[i-j : i+j+1]
				}
			} else {
				break
			}
		}
	}
	for i := 1; i < len(s); i++ { // case "abba"
		for j := 0; i-j-1 >= 0 && i+j < len(s); j++ {
			if s[i-j-1] == s[i+j] {
				if len(result) < 2*j+2 {
					result = s[i-j-1 : i+j+1]
				}
			} else {
				break
			}
		}
	}
	return result
}
