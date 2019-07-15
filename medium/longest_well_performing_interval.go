package medium

// We are given hours, a list of the number of hours worked
// per day for a given employee.
//
// A day is considered to be a tiring day if and only if the
// number of hours worked is (strictly) greater than 8.
//
// A well-performing interval is an interval of days for which
// the number of tiring days is strictly larger than the number
// of non-tiring days.
//
// Return the length of the longest well-performing interval.
//
// Example 1:
// Input: hours = [9,9,6,0,6,6,9]
// Output: 3
// Explanation: The longest well-performing interval is [9,9,6].
//
// Constraints:
// 1 <= hours.length <= 10000
// 0 <= hours[i] <= 16

func longestWPI(hours []int) int {
	var dp []int
	var cnt int
	for i := 0; i < len(hours); i++ {
		if hours[i] > 8 {
			cnt++
		}
		dp = append(dp, cnt)
	}

	var max int
	for i := 0; i < len(hours); i++ {
		for j := i; j < len(hours); j++ {
			var before int
			if i-1 < 0 {
				before = 0
			} else {
				before = dp[i-1]
			}
			if dp[j]-before > (j-i+1)/2 && max < j-i+1 {
				max = j - i + 1
			}
		}
	}
	return max
}

// My Time Limit Exceeded Version:
//
// func longestWPI(hours []int) int {
// 	var max int
// 	for i := 0; i < len(hours); i++ {
// 		for j := i; j < len(hours); j++ {
// 			if isValidHours(hours, i, j) && max < j-i+1 {
// 				max = j - i + 1
// 			}
// 		}
// 	}
// 	return max
// }
//
// func isValidHours(hours []int, i, j int) bool {
// 	var n int
// 	for k := i; k <= j; k++ {
// 		if hours[k] > 8 {
// 			n++
// 		}
// 	}
// 	if n > (j-i+1)/2 {
// 		return true
// 	}
// 	return false
// }
