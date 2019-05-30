package main

func addDigits(num int) int {
	ans := 0
	for num != 0 {
		ans += num % 10
		num /= 10
	}
	if ans < 10 {
		return ans
	}
	return addDigits(ans)
}
