package easy

// Given a list of dominoes, dominoes[i] = [a, b] is equivalent to
// dominoes[j] = [c, d] if and only if either (a==c and b==d), or
// (a==d and b==c) - that is, one domino can be rotated to be equal
// to another domino.
//
// Return the number of pairs (i, j) for which 0 <= i < j < dominoes.length,
// and dominoes[i] is equivalent to dominoes[j].
//
// Example 1:
// Input: dominoes = [[1,2],[2,1],[3,4],[5,6]]
// Output: 1
//
// Constraints:
// 1 <= dominoes.length <= 40000
// 1 <= dominoes[i][j] <= 9

func numEquivDominoPairs(dominoes [][]int) int {
	cnt := 0
	for i := 0; i < len(dominoes); i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if isDomino(dominoes[i], dominoes[j]) {
				cnt++
			}
		}
	}
	return cnt
}

func isDomino(a, b []int) bool {
	if len(a) != 2 || len(b) != 2 {
		return false
	}
	if a[0] == b[0] && a[1] == b[1] {
		return true
	}
	if a[0] == b[1] && a[1] == b[0] {
		return true
	}
	return false
}

// Answers of other people of contest:
//
// class Solution {
// 	public:
// 		int numEquivDominoPairs(vector<vector<int>>& dominoes) {
// 			map<pair<int, int>, int> freq;
// 			long long total = 0;
//
// 			for (vector<int> domino : dominoes) {
// 				if (domino[0] > domino[1])
// 					swap(domino[0], domino[1]);
//
// 				total += freq[make_pair(domino[0], domino[1])]++;
// 			}
//
// 			return total;
// 		}
// 	};
//
// Answer of other people in discuss:
//
// public int numEquivDominoPairs(int[][] dominoes) {
// 	Map<Integer, Integer> count = new HashMap<>();
// 	int res = 0;
// 	for (int[] d : dominoes) {
// 		int k = Math.min(d[0], d[1]) * 10 + Math.max(d[0], d[1]);
// 		count.put(k, count.getOrDefault(k, 0) + 1);
// 	}
// 	for (int v : count.values()) {
// 		res += v * (v - 1) / 2;
// 	}
// 	return res;
// }
