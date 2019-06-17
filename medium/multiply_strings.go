package medium

// Given two non-negative integers num1 and num2 represented as strings,
// return the product of num1 and num2, also represented as a string.
//
// Example 1:
// Input: num1 = "2", num2 = "3"
// Output: "6"
//
// Example 2:
// Input: num1 = "123", num2 = "456"
// Output: "56088"
//
// Note:
// The length of both num1 and num2 is < 110.
// Both num1 and num2 contain only digits 0-9.
// Both num1 and num2 do not contain any leading zero, except the number 0 itself.
// You must not use any built-in BigInteger library or convert the inputs to integer directly.

func multiply(num1 string, num2 string) string {
	if len(num1) == 0 || len(num2) == 0 {
		return ""
	}

	ans := make([]byte, len(num1)+len(num2))
	carry := func(i int) {
		for ans[i] >= 10 {
			ans[i+1] += ans[i] / 10
			ans[i] %= 10
			i++
		}
	}

	num1 = reversal(num1)
	num2 = reversal(num2)
	for i := 0; i < len(num1); i++ {
		for j := 0; j < len(num2); j++ {
			v := (num2[j] - '0') * (num1[i] - '0')
			ans[i+j] += v % 10
			carry(i + j)
			ans[i+j+1] += v / 10
			carry(i + j + 1)
		}
	}

	noZeroIdx := -1
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] != 0 {
			noZeroIdx = i
			break
		}
	}
	if noZeroIdx == -1 {
		return "0"
	}

	ans = ans[:noZeroIdx+1]
	for i := 0; i < len(ans); i++ {
		ans[i] += '0'
	}

	return reversal(string(ans))
}

func reversal(s string) string {
	var ans []byte
	for i := len(s) - 1; i >= 0; i-- {
		ans = append(ans, s[i])
	}
	return string(ans)
}
