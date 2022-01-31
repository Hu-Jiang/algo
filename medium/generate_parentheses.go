package medium

import (
	"container/list"
	"errors"
)

// Given n pairs of parentheses, write a function to generate all combinations
// of well-formed parentheses.
//
// For example, given n = 3, a solution set is:
// [
//   "((()))",
//   "(()())",
//   "(())()",
//   "()(())",
//   "()()()"
// ]

func generateParenthesis(n int) []string {
	if n <= 0 {
		return nil
	}

	var combi []string
	genarateAllCombination(2*n, "", &combi)

	var ans []string
	for i := 0; i < len(combi); i++ {
		if isValid(combi[i]) {
			ans = append(ans, combi[i])
		}
	}

	return ans
}

func isValid(s string) bool {
	q := newQueue()
	valid := true

	for _, r := range s {
		switch r {
		case '(':
			q.enqueue(r)
		case ')':
			v, err := q.dequeue()
			if err != nil || v != '(' {
				valid = false
				break
			}

		}
	}

	return valid && q.isEmpty()
}

func genarateAllCombination(n int, prefix string, res *[]string) {
	if n == 0 {
		*res = append(*res, prefix)
		return
	}

	genarateAllCombination(n-1, prefix+"(", res)
	genarateAllCombination(n-1, prefix+")", res)
}

type queue struct {
	l *list.List
}

func newQueue() *queue {
	return &queue{l: list.New()}
}

func (q *queue) enqueue(v interface{}) {
	q.l.PushBack(v)
}

func (q *queue) isEmpty() bool {
	return q.l.Len() == 0
}

func (q *queue) dequeue() (interface{}, error) {
	if q.isEmpty() {
		return nil, errors.New("dequeue failed cause of queue is empty")
	}
	return q.l.Remove(q.l.Back()), nil
}

// The official solution:
//
// Approach 1: Brute Force
//
// class Solution {
//     public List<String> generateParenthesis(int n) {
//         List<String> combinations = new ArrayList();
//         generateAll(new char[2 * n], 0, combinations);
//         return combinations;
//     }
//
//     public void generateAll(char[] current, int pos, List<String> result) {
//         if (pos == current.length) {
//             if (valid(current))
//                 result.add(new String(current));
//         } else {
//             current[pos] = '(';
//             generateAll(current, pos+1, result);
//             current[pos] = ')';
//             generateAll(current, pos+1, result);
//         }
//     }
//
//     public boolean valid(char[] current) {
//         int balance = 0;
//         for (char c: current) {
//             if (c == '(') balance++;
//             else balance--;
//             if (balance < 0) return false;
//         }
//         return (balance == 0);
//     }
// }
//
// Approach 2: Backtracking
//
// class Solution {
//     public List<String> generateParenthesis(int n) {
//         List<String> ans = new ArrayList();
//         backtrack(ans, "", 0, 0, n);
//         return ans;
//     }
//
//     public void backtrack(List<String> ans, String cur, int open, int close, int max){
//         if (cur.length() == max * 2) {
//             ans.add(cur);
//             return;
//         }
//
//         if (open < max)
//             backtrack(ans, cur+"(", open+1, close, max);
//         if (close < open)
//             backtrack(ans, cur+")", open, close+1, max);
//     }
// }
//
// Approach 3: Closure Number
//
