package medium

// There are n flights, and they are labeled from 1 to n.
//
// We have a list of flight bookings.  The i-th booking
// bookings[i] = [i, j, k] means that we booked k seats
// from flights labeled i to j inclusive.
//
// Return an array answer of length n, representing the
// number of seats booked on each flight in order of
// their label.
//
// Example 1:
// Input: bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
// Output: [10,55,45,25,25]
//
// Constraints:
// 1 <= bookings.length <= 20000
// 1 <= bookings[i][0] <= bookings[i][1] <= n <= 20000
// 1 <= bookings[i][2] <= 10000

func corpFlightBookings(bookings [][]int, n int) []int {
	if len(bookings) == 0 {
		return nil
	}

	res := make([]int, n+1)
	diff := make([]int, n+2)
	for i := 0; i < len(bookings); i++ {
		diff[bookings[i][0]] += bookings[i][2]
		diff[bookings[i][1]+1] -= bookings[i][2]
	}
	for i := 1; i <= n; i++ {
		res[i] += res[i-1] + diff[i]
	}

	return res[1:]
}

// My First Version:
//
// func corpFlightBookings(bookings [][]int, n int) []int {
// 	if len(bookings) == 0 {
// 		return nil
// 	}
//
// 	res := make([]int, n+1)
// 	for i := 0; i < len(bookings); i++ {
// 		for j := bookings[i][0]; j <= bookings[i][1]; j++ {
// 			res[j] += bookings[i][2]
// 		}
// 	}
// 	return res[1:]
// }
