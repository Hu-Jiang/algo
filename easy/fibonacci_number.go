package easy

// The Fibonacci numbers, commonly denoted F(n) form a sequence,
// called the Fibonacci sequence, such that each number is the
// sum of the two preceding ones, starting from 0 and 1. That is,
// F(0) = 0,   F(1) = 1
// F(N) = F(N - 1) + F(N - 2), for N > 1.
//
// Given N, calculate F(N).
//
// Example 1:
// Input: 2
// Output: 1
// Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.
//
// Example 2:
// Input: 3
// Output: 2
// Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.
//
// Example 3:
// Input: 4
// Output: 3
// Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
//
// Note:
// 0 ≤ N ≤ 30.

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	first, second := 0, 1
	cur := 0
	for i := 2; i <= n; i++ {
		cur = first + second
		first, second = second, cur
	}

	return cur
}

// The official soloution:
//
// Approach 1: Recursion
//
// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     return fib(N-1) + fib(N-2)
// }
//
// Approach 2: Bottom-Up Approach using Memoization
//
// class Solution {
//     public int fib(int N) {
//         if (N <= 1) {
//             return N;
//         }
//         return memoize(N);
//     }
//
//     public int memoize(int N) {
//       int[] cache = new int[N + 1];
//       cache[1] = 1;
//
//       for (int i = 2; i <= N; i++) {
//           cache[i] = cache[i-1] + cache[i-2];
//       }
//       return cache[N];
//     }
// }
//
// Approach 3: Top-Down Approach using Memoization
//
// class Solution {
//     private Integer[] cache = new Integer[31];
//
//     public int fib(int N) {
//         if (N <= 1) {
//             return N;
//         }
//         cache[0] = 0;
//         cache[1] = 1;
//         return memoize(N);
//     }
//
//     public int memoize(int N) {
//       if (cache[N] != null) {
//           return cache[N];
//       }
//       cache[N] = memoize(N-1) + memoize(N-2);
//       return memoize(N);
//     }
// }
//
// Approach 4: Iterative Top-Down Approach
//
// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     if N == 2 {
//         return 1
//     }
//
//     current := 0
//     prev1 := 1
//     prev2 := 1
//
//     for i := 3; i <= N; i++ {
//         current = prev1 + prev2
//         prev2 = prev1
//         prev1 = current
//     }
//     return current
// }
//
// Approach 5: Matrix Exponentiation
//
// func fib(N int) int {
//     if N <= 1 {
//         return N
//     }
//     var A = [2][2]int{
//         {1,1},
//         {1,0},
//     }
//     A = matrixPower(A, N-1)
//     return A[0][0]
// }
//
// func matrixPower(A [2][2]int, N int) [2][2]int {
//     if N <= 1 {
//         return A
//     }
//     A = matrixPower(A, N/2)
//     A = multiply(A, A)
//
//     var B = [2][2]int{
//         {1,1},
//         {1,0},
//     }
//     if N%2 != 0 {
//         A = multiply(A, B)
//     }
//
//     return A
// }
//
// func multiply(A [2][2]int, B [2][2]int) [2][2]int {
//     x := A[0][0] * B[0][0] + A[0][1] * B[1][0];
//     y := A[0][0] * B[0][1] + A[0][1] * B[1][1];
//     z := A[1][0] * B[0][0] + A[1][1] * B[1][0];
//     w := A[1][0] * B[0][1] + A[1][1] * B[1][1];
//
//     A[0][0] = x;
//     A[0][1] = y;
//     A[1][0] = z;
//     A[1][1] = w;
//
//     return A
// }
//
// Approach 6: Math
//
// func fib(N int) int {
//     var goldenRatio float64 = float64((1 + math.Sqrt(5)) / 2);
//     return int(math.Round(math.Pow(goldenRatio, float64(N)) / math.Sqrt(5)));
// }
