package medium

// On an alphabet board, we start at position (0, 0), corresponding
// to character board[0][0].
//
// Here, board = ["abcde", "fghij", "klmno", "pqrst", "uvwxy", "z"],
// as shown in the diagram below.
// Image location: [/image/medium/alphabet_board_path.png]
//
// We may make the following moves:
// 'U' moves our position up one row, if the position exists on the board;
// 'D' moves our position down one row, if the position exists on the board;
// 'L' moves our position left one column, if the position exists on the board;
// 'R' moves our position right one column, if the position exists on the board;
// '!' adds the character board[r][c] at our current position (r, c) to the answer.
// (Here, the only positions that exist on the board are positions with letters on them.)
//
// Return a sequence of moves that makes our answer equal to target in the minimum number
// of moves. You may return any path that does so.
//
// Example 1:
// Input: target = "leet"
// Output: "DDR!UURRR!!DDD!"
//
// Example 2:
// Input: target = "code"
// Output: "RR!DDRR!UUL!R!"
//
// Constraints:
// 1 <= target.length <= 100
// target consists only of English lowercase letters.

type position struct {
	x int
	y int
}

func alphabetBoardPath(target string) string {
	runePosition := make(map[rune]position)
	for i := 0; i < 26; i++ {
		runePosition[rune('a'+i)] = position{i / 5, i % 5}
	}

	var result string
	prevPosition := runePosition['a']
	for _, r := range target {
		diffx := runePosition[r].x - prevPosition.x
		diffy := runePosition[r].y - prevPosition.y

		if diffx < 0 {
			for i := 0; i < -diffx; i++ {
				result += "U"
			}
		}
		if diffy < 0 {
			for i := 0; i < -diffy; i++ {
				result += "L"
			}
		}
		if diffx > 0 {
			for i := 0; i < diffx; i++ {
				result += "D"
			}
		}
		if diffy > 0 {
			for i := 0; i < diffy; i++ {
				result += "R"
			}
		}

		result += "!"
		prevPosition = runePosition[r]
	}

	return result
}

// Another answer from discuss in leetcode:
//
// def alphabetBoardPath(self, target):
// m = {c: [i / 5, i % 5] for i, c in enumerate("abcdefghijklmnopqrstuvwxyz")}
// x0, y0 = 0, 0
// res = []
// for c in target:
// 	x, y = m[c]
// 	if y < y0: res.append('L' * (y0 - y))
// 	if x < x0: res.append('U' * (x0 - x))
// 	if x > x0: res.append('D' * (x - x0))
// 	if y > y0: res.append('R' * (y - y0))
// 	res.append('!')
// 	x0, y0 = x, y
// return "".join(res)
