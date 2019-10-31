package hard

import (
	"container/list"
	"errors"
	"strings"
)

// Given a string representing a code snippet, you need to implement a tag validator
// to parse the code and return whether it is valid. A code snippet is valid if all
// the following rules hold:
//
// 1. The code must be wrapped in a valid closed tag. Otherwise, the code is invalid.
// 2. A closed tag (not necessarily valid) has exactly the following format :
//    <TAG_NAME>TAG_CONTENT</TAG_NAME>. Among them, <TAG_NAME> is the start tag,
//    and </TAG_NAME> is the end tag. The TAG_NAME in start and end tags should
//    be the same. A closed tag is valid if and only if the TAG_NAME and
//    TAG_CONTENT are valid.
// 3. A valid TAG_NAME only contain upper-case letters, and has length in range [1,9].
// 	  Otherwise, the TAG_NAME is invalid.
// 4. A valid TAG_CONTENT may contain other valid closed tags, cdata and any characters
// 	  (see note1) EXCEPT unmatched <, unmatched start and end tag, and unmatched or
//    closed tags with invalid TAG_NAME. Otherwise, the TAG_CONTENT is invalid.
// 5. A start tag is unmatched if no end tag exists with the same TAG_NAME, and vice versa.
//    However, you also need to consider the issue of unbalanced when tags are nested.
// 6. A < is unmatched if you cannot find a subsequent >. And when you find a < or </,
//    all the subsequent characters until the next > should be parsed as
//    TAG_NAME (not necessarily valid).
// 7. The cdata has the following format : <![CDATA[CDATA_CONTENT]]>. The range of
//    CDATA_CONTENT is defined as the characters between <![CDATA[ and the first subsequent ]]>.
// 8. CDATA_CONTENT may contain any characters. The function of cdata is to forbid the validator
//    to parse CDATA_CONTENT, so even it has some characters that can be parsed as tag
//    (no matter valid or invalid), you should treat it as regular characters.
//
// Valid Code Examples:
//
// Input: "<DIV>This is the first line <![CDATA[<div>]]></DIV>"
// Output: True
// Explanation:
// The code is wrapped in a closed tag : <DIV> and </DIV>.
// The TAG_NAME is valid, the TAG_CONTENT consists of some characters and cdata.
// Although CDATA_CONTENT has unmatched start tag with invalid TAG_NAME, it should
// be considered as plain text, not parsed as tag.
// So TAG_CONTENT is valid, and then the code is valid. Thus return true.
//
// Input: "<DIV>>>  ![cdata[]] <![CDATA[<div>]>]]>]]>>]</DIV>"
// Output: True
// Explanation:
// We first separate the code into : start_tag|tag_content|end_tag.
// start_tag -> "<DIV>"
// end_tag -> "</DIV>"
// tag_content could also be separated into : text1|cdata|text2.
// text1 -> ">>  ![cdata[]] "
// cdata -> "<![CDATA[<div>]>]]>", where the CDATA_CONTENT is "<div>]>"
// text2 -> "]]>>]"
// The reason why start_tag is NOT "<DIV>>>" is because of the rule 6.
// The reason why cdata is NOT "<![CDATA[<div>]>]]>]]>" is because of the rule 7.
//
// Invalid Code Examples:
//
// Input: "<A>  <B> </A>   </B>"
// Output: False
// Explanation: Unbalanced. If "<A>" is closed, then "<B>" must be unmatched, and vice versa.
//
// Input: "<DIV>  div tag is not closed  <DIV>"
// Output: False
//
// Input: "<DIV>  unmatched <  </DIV>"
// Output: False
//
// Input: "<DIV> closed tags with invalid tag name  <b>123</b> </DIV>"
// Output: False
//
// Input: "<DIV> unmatched tags with invalid tag name  </1234567890> and <CDATA[[]]>  </DIV>"
// Output: False
//
// Input: "<DIV>  unmatched start tag <B>  and unmatched end tag </C>  </DIV>"
// Output: False
//
// Note:
// For simplicity, you could assume the input code (including the any characters mentioned above)
// only contain letters, digits, '<','>','/','!','[',']' and ' '.

func isValid(code string) bool {
	if len(code) == 0 {
		return false
	}

	if code[0] != '<' || code[len(code)-1] != '>' {
		return false
	}

	ts := newStack()
	i := 0
	for i < len(code) {
		if i != 0 && ts.isEmpty() {
			return false
		}
		if code[i] == '<' {
			switch code[i+1] { // i+1 must be valid, because last item is '>'
			case '!':
				if i == 0 {
					return false
				}
				cdata, err := getCDATA(code[i:])
				if err != nil {
					return false
				}
				i += len(cdata)
			case '/':
				etag, err := getEndTagName(code[i:])
				if err != nil {
					return false
				}
				if !etag.isValid() {
					return false
				}

				v, err := ts.pop()
				if err != nil {
					return false
				}
				if v.(tagName) != etag {
					return false
				}
				i += len(etag) + 3
			default:
				stag, err := getStartTagName(code[i:])
				if err != nil {
					return false
				}
				if !stag.isValid() {
					return false
				}

				ts.push(stag)
				i += len(stag) + 2
			}
		} else {
			i++
		}
	}

	return ts.isEmpty()
}

// A tagName represents a TAG_NAME sotred in stack.
type tagName string

func (t tagName) isValid() bool {
	if len(t) < 1 || len(t) > 9 {
		return false
	}

	for _, v := range t {
		if v < 'A' || v > 'Z' {
			return false
		}
	}
	return true
}

func getStartTagName(code string) (tagName, error) {
	i1 := strings.IndexRune(code, '<')
	i2 := strings.IndexRune(code, '>')
	if i1 > i2 || i1 == -1 || i2 == -1 {
		return "", errors.New("invalid start tag symbol")
	}

	return tagName(code[i1+1 : i2]), nil

}

func getEndTagName(code string) (tagName, error) {
	i1 := strings.Index(code, "</")
	i2 := strings.IndexRune(code, '>')
	if i1 > i2 || i1 == -1 || i2 == -1 {
		return "", errors.New("invalid end tag symbol")
	}

	return tagName(code[i1+2 : i2]), nil
}

func getCDATA(code string) (string, error) {
	i1 := strings.Index(code, "<![CDATA[")
	i2 := strings.Index(code, "]]>")
	if i1 > i2 || i1 == -1 || i2 == -1 {
		return "", errors.New("invalid cdata symbol")
	}
	return code[i1 : i2+3], nil
}

type stack struct {
	s *list.List
}

func newStack() *stack {
	return &stack{s: list.New()}
}

func (s *stack) push(v interface{}) {
	s.s.PushBack(v)
}

func (s *stack) isEmpty() bool {
	return s.s.Len() == 0
}

func (s *stack) pop() (interface{}, error) {
	if s.isEmpty() {
		return nil, errors.New("pop: empty stack")
	}
	return s.s.Remove(s.s.Back()), nil
}

// The official solution:
//
// Approach 1: Stack
//
// public class Solution {
//     Stack < String > stack = new Stack < > ();
//     boolean contains_tag = false;
//     public boolean isValidTagName(String s, boolean ending) {
//         if (s.length() < 1 || s.length() > 9)
//             return false;
//         for (int i = 0; i < s.length(); i++) {
//             if (!Character.isUpperCase(s.charAt(i)))
//                 return false;
//         }
//         if (ending) {
//             if (!stack.isEmpty() && stack.peek().equals(s))
//                 stack.pop();
//             else
//                 return false;
//         } else {
//             contains_tag = true;
//             stack.push(s);
//         }
//         return true;
//     }
//     public boolean isValidCdata(String s) {
//         return s.indexOf("[CDATA[") == 0;
//     }
//     public boolean isValid(String code) {
//         if (code.charAt(0) != '<' || code.charAt(code.length() - 1) != '>')
//             return false;
//         for (int i = 0; i < code.length(); i++) {
//             boolean ending = false;
//             int closeindex;
//             if(stack.isEmpty() && contains_tag)
//                 return false;
//             if (code.charAt(i) == '<') {
//                 if (!stack.isEmpty() && code.charAt(i + 1) == '!') {
//                     closeindex = code.indexOf("]]>", i + 1);
//                     if (closeindex < 0 || !isValidCdata(code.substring(i + 2, closeindex)))
//                         return false;
//                 } else {
//                     if (code.charAt(i + 1) == '/') {
//                         i++;
//                         ending = true;
//                     }
//                     closeindex = code.indexOf('>', i + 1);
//                     if (closeindex < 0 || !isValidTagName(code.substring(i + 1, closeindex), ending))
//                         return false;
//                 }
//                 i = closeindex;
//             }
//         }
//         return stack.isEmpty() && contains_tag;
//     }
// }
//
// Approach 2: Regex
//
// import java.util.regex.*;
// public class Solution {
//     Stack < String > stack = new Stack < > ();
//     boolean contains_tag = false;
//     public boolean isValidTagName(String s, boolean ending) {
//         if (ending) {
//             if (!stack.isEmpty() && stack.peek().equals(s))
//                 stack.pop();
//             else
//                 return false;
//         } else {
//             contains_tag = true;
//             stack.push(s);
//         }
//         return true;
//     }
//     public boolean isValid(String code) {
//         String regex = "<[A-Z]{0,9}>([^<]*(<((\\/?[A-Z]{1,9}>)|(!\\[CDATA\\[(.*?)]]>)))?)*";
//         if (!Pattern.matches(regex, code))
//             return false;
//         for (int i = 0; i < code.length(); i++) {
//             boolean ending = false;
//             if (stack.isEmpty() && contains_tag)
//                 return false;
//             if (code.charAt(i) == '<') {
//                 if (code.charAt(i + 1) == '!') {
//                     i = code.indexOf("]]>", i + 1);
//                     continue;
//                 }
//                 if (code.charAt(i + 1) == '/') {
//                     i++;
//                     ending = true;
//                 }
//                 int closeindex = code.indexOf('>', i + 1);
//                 if (closeindex < 0 || !isValidTagName(code.substring(i + 1, closeindex), ending))
//                     return false;
//                 i = closeindex;
//             }
//         }
//         return stack.isEmpty();
//     }
// }
