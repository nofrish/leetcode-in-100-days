package _012

import "strings"

func intToRoman(num int) string {

	mapping := [][]string{
		{"I", "IV", "V", "IX"},
		{"X", "XL", "L", "XC"},
		{"C", "CD", "D", "CM"},
		{"M"},
	}

	result := ""
	for row := 0; num > 0; row, num = row+1, num/10 {
		r, sb := num%10, strings.Builder{}
		switch {
		case r <= 3: // 第一种情况
			for i := 0; i < r; i++ {
				sb.WriteString(mapping[row][0])
			}
		case r == 4: // 第二种情况
			sb.WriteString(mapping[row][1])
		case r > 4 && r < 9: // 第三种情况
			sb.WriteString(mapping[row][2])
			for i := 5; i < r; i++ {
				sb.WriteString(mapping[row][0])
			}
		case r == 9: // 第四种情况
			sb.WriteString(mapping[row][3])
		}
		result = sb.String() + result
	}
	return result
}
