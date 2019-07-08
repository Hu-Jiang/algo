package hard

// Given an array of words and a width maxWidth, format the text such that
// each line has exactly maxWidth characters and is fully (left and right) justified.
//
// You should pack your words in a greedy approach; that is, pack as many words as you
// can in each line. Pad extra spaces ' ' when necessary so that each line has exactly
// maxWidth characters.
//
// Extra spaces between words should be distributed as evenly as possible. If the number
// of spaces on a line do not divide evenly between words, the empty slots on the left
// will be assigned more spaces than the slots on the right.
//
// For the last line of text, it should be left justified and no extra space is inserted
// between words.
//
// Note:
// A word is defined as a character sequence consisting of non-space characters only.
// Each word's length is guaranteed to be greater than 0 and not exceed maxWidth.
// The input array words contains at least one word.
//
// Example 1:
// Input:
// words = ["This", "is", "an", "example", "of", "text", "justification."]
// maxWidth = 16
// Output:
// [
//    "This    is    an",
//    "example  of text",
//    "justification.  "
// ]
//
// Example 2:
// Input:
// words = ["What","must","be","acknowledgment","shall","be"]
// maxWidth = 16
// Output:
// [
//   "What   must   be",
//   "acknowledgment  ",
//   "shall be        "
// ]
// Explanation: Note that the last line is "shall be    " instead of "shall     be",
//              because the last line must be left-justified instead of fully-justified.
//              Note that the second line is also left-justified becase it contains only one word.
//
// Example 3:
// Input:
// words = ["Science","is","what","we","understand","well","enough","to","explain",
//          "to","a","computer.","Art","is","everything","else","we","do"]
// maxWidth = 20
// Output:
// [
//   "Science  is  what we",
//   "understand      well",
//   "enough to explain to",
//   "a  computer.  Art is",
//   "everything  else  we",
//   "do                  "
// ]

func fullJustify(words []string, maxWidth int) []string {
	if len(words) == 0 {
		return []string{genSpace(maxWidth)}
	}
	var res []string
	for i := 0; i < len(words); i++ {
		wordWidth, curLen, j := 0, 0, i
		for ; j < len(words) && curLen <= maxWidth; j++ {
			wordWidth += len(words[j])
			curLen += len(words[j])
			if j != i {
				curLen += 1
			}
		}
		if j >= len(words) && curLen <= maxWidth {
			j--
		} else {
			j -= 2
		}

		// in here, words[i, j] is the new line

		if curLen > maxWidth {
			wordWidth -= len(words[j+1])
		}

		var tmp string
		space := maxWidth - wordWidth
		var base, more int
		if j-i != 0 {
			base = space / (j - i)
			more = space - base*(j-i)
		}

		for k := i; k <= j; k++ {
			appendSpace := 1
			if space == 0 {
				appendSpace = 0
			}
			if j != len(words)-1 && j != k {
				if j-k == 1 {
					appendSpace = space
				} else {
					if more != 0 {
						appendSpace = base + 1
						more--
					} else {
						appendSpace = base
					}
				}
			}

			tmp += words[k] + genSpace(appendSpace)
			space -= appendSpace
		}
		if len(tmp) < maxWidth {
			tmp += genSpace(space)
		}
		res = append(res, tmp)
		i = j
	}

	return res
}

func genSpace(n int) string {
	var res string
	for i := 0; i < n; i++ {
		res += " "
	}
	return res
}
