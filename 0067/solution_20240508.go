package _067

func addBinary(a string, b string) string {

	result := make([]byte, max(len(a), len(b))+1)

	i, j, cur := len(a)-1, len(b)-1, len(result)-1
	carry := byte(0)
	for ; i >= 0 && j >= 0; i, j, cur = i-1, j-1, cur-1 {
		sum := carry + (a[i] - '0') + (b[j] - '0')
		result[cur] = '0' + sum%2
		carry = sum / 2
	}

	for ; i >= 0; i, cur = i-1, cur-1 {
		sum := carry + (a[i] - '0')
		result[cur] = '0' + sum%2
		carry = sum / 2
	}

	for ; j >= 0; j, cur = j-1, cur-1 {
		sum := carry + (b[j] - '0')
		result[cur] = '0' + sum%2
		carry = sum / 2
	}

	if carry == 1 {
		result[0] = '1'
		return string(result)
	}
	return string(result[1:])
}
