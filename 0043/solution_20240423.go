package _043

import "strings"

// 基本思路是：先考虑计算每一位的乘法结果，再考虑进位
// 基本方法如下：
//
//               9   9   9
//         ×     6   7   8
//  ----------------------
//              72  72  72
//          63  63  63
//      54  54  54
//  ----------------------
//      54 117 189 135  72
//  ----------------------
//      54 117 189 142   2
//  -----------------------
//      54 117 203   2   2
//  -----------------------
//      54 137   3   2   2
//  -----------------------
//      67   7   3   2   2
//  -----------------------
//   6   7   7   3   2   2

func multiply(num1 string, num2 string) string {

	if num1 == "0" || num2 == "0" { // 当一个乘数为"0"的时候，直接返回"0"
		return "0"
	}

	product := make([]int, len(num1)+len(num2)) // 最终的结果存储在一个 len(num1)+len(num2) 的数组中
	for i := len(num2) - 1; i >= 0; i-- {
		for j := len(num1) - 1; j >= 0; j-- {
			product[i+j+1] += int((num2[i] - '0') * (num1[j] - '0')) // 将结果存入到对应的位置，不考虑进位
		}
	}
	for i, r := len(product)-1, 0; i >= 0; i-- {
		product[i], r = (product[i]+r)%10, (product[i]+r)/10 // 进位
	}

	if product[0] == 0 {
		product = product[1:] // 处理开头的"0"
	}

	sb := strings.Builder{}
	for i := 0; i < len(product); i++ {
		sb.WriteByte('0' + byte(product[i]))
	}
	return sb.String()
}
