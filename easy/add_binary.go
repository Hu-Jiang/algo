package easy

// Given two binary strings, return their sum (also a binary string).
//
// The input strings are both non-empty and contains only characters 1 or 0.
//
// Example 1:
// Input: a = "11", b = "1"
// Output: "100"
//
// Example 2:
// Input: a = "1010", b = "1011"
// Output: "10101"

func addBinary(a string, b string) string {
	bytea, byteb := convert(a), convert(b)
	newByte := make([]byte, maxVal(len(bytea), len(byteb)))
	var carry byte = 0
	i, j, k := len(bytea)-1, len(byteb)-1, len(newByte)-1
	var sum byte
	for i >= 0 && j >= 0 {
		sum = bytea[i] + byteb[j] + carry
		newByte[k] = sum % 2
		carry = sum / 2
		i--
		j--
		k--
	}

	for i >= 0 {
		sum = bytea[i] + carry
		newByte[k] = sum % 2
		carry = sum / 2
		i--
		k--
	}

	for j >= 0 {
		sum = byteb[j] + carry
		newByte[k] = sum % 2
		carry = sum / 2
		j--
		k--
	}

	for i := 0; i < len(newByte); i++ {
		newByte[i] += '0'
	}

	if carry != 0 {
		return "1" + string(newByte)
	}

	return string(newByte)

}

func convert(str string) []byte {
	b := make([]byte, len(str))
	for i := 0; i < len(b); i++ {
		b[i] = str[i] - '0'
	}
	return b
}

func maxVal(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
