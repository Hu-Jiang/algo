package medium

import (
	"strconv"
)

// Given a string containing digits from 2-9 inclusive, return all possible letter
// combinations that the number could represent.
//
// A mapping of digit to letters (just like on the telephone buttons) is given below.
// Note that 1 does not map to any letters.
//
// Image location: [/image/medium/letter_combinations_of_a_phone_number.png]
//
// Example:
// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
// Note:
// Although the above answer is in lexicographical order, your answer could be
// in any order you want.

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var ans []string
	nextAppend("", digits, &ans)
	return ans
}

var letters = [...]string{"", "", "abc", "def", "ghi", "jkl",
	"mno", "pqrs", "tuv", "wxyz"}

func nextAppend(prefix string, remainDigits string, ans *[]string) {
	if len(remainDigits) == 0 {
		*ans = append(*ans, prefix)
		return
	}
	v, _ := strconv.Atoi(string(remainDigits[0]))
	for i := 0; i < len(letters[v]); i++ {
		nextAppend(prefix+string(letters[v][i]), remainDigits[1:], ans)
	}
}

// The official soloution:
//
// Approach 1: Backtracking
//
// class Solution {
// 	Map<String, String> phone = new HashMap<String, String>() {{
// 	  put("2", "abc");
// 	  put("3", "def");
// 	  put("4", "ghi");
// 	  put("5", "jkl");
// 	  put("6", "mno");
// 	  put("7", "pqrs");
// 	  put("8", "tuv");
// 	  put("9", "wxyz");
// 	}};
//
// 	List<String> output = new ArrayList<String>();
//
// 	public void backtrack(String combination, String next_digits) {
// 	  // if there is no more digits to check
// 	  if (next_digits.length() == 0) {
// 		// the combination is done
// 		output.add(combination);
// 	  }
// 	  // if there are still digits to check
// 	  else {
// 		// iterate over all letters which map
// 		// the next available digit
// 		String digit = next_digits.substring(0, 1);
// 		String letters = phone.get(digit);
// 		for (int i = 0; i < letters.length(); i++) {
// 		  String letter = phone.get(digit).substring(i, i + 1);
// 		  // append the current letter to the combination
// 		  // and proceed to the next digits
// 		  backtrack(combination + letter, next_digits.substring(1));
// 		}
// 	  }
// 	}
//
// 	public List<String> letterCombinations(String digits) {
// 	  if (digits.length() != 0)
// 		backtrack("", digits);
// 	  return output;
// 	}
// }
