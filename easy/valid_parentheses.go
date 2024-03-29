package easy

import (
	"container/list"
	"errors"
)

// Given a string containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Note that an empty string is also considered valid.
//
// Example 1:
// Input: "()"
// Output: true
//
// Example 2:
// Input: "()[]{}"
// Output: true
//
// Example 3:
// Input: "(]"
// Output: false
//
// Example 4:
// Input: "([)]"
// Output: false
//
// Example 5:
// Input: "{[]}"
// Output: true

func isValid(s string) bool {
	q := newQueue()
	valid := true

	for _, r := range s {
		switch r {
		case '(', '{', '[':
			q.enqueue(r)
		case ')', '}', ']':
			if q.isEmpty() {
				valid = false
				break
			} else {
				v, _ := q.dequeue()
				if bracketsPair[v.(rune)] != r {
					valid = false
					break
				}
			}
		}
	}

	return valid && q.isEmpty()
}

var bracketsPair = map[rune]rune{
	'(': ')',
	'{': '}',
	'[': ']',
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
// Approach 1: Stacks
//
// class Solution {
//
// 	// Hash table that takes care of the mappings.
// 	private HashMap<Character, Character> mappings;
//
// 	// Initialize hash map with mappings. This simply makes the code easier to read.
// 	public Solution() {
// 	  this.mappings = new HashMap<Character, Character>();
// 	  this.mappings.put(')', '(');
// 	  this.mappings.put('}', '{');
// 	  this.mappings.put(']', '[');
// 	}
//
// 	public boolean isValid(String s) {
//
// 	  // Initialize a stack to be used in the algorithm.
// 	  Stack<Character> stack = new Stack<Character>();
//
// 	  for (int i = 0; i < s.length(); i++) {
// 		char c = s.charAt(i);
//
// 		// If the current character is a closing bracket.
// 		if (this.mappings.containsKey(c)) {
//
// 		  // Get the top element of the stack. If the stack is empty, set a dummy value of '#'
// 		  char topElement = stack.empty() ? '#' : stack.pop();
//
// 		  // If the mapping for this bracket doesn't match the stack's top element, return false.
// 		  if (topElement != this.mappings.get(c)) {
// 			return false;
// 		  }
// 		} else {
// 		  // If it was an opening bracket, push to the stack.
// 		  stack.push(c);
// 		}
// 	  }
//
// 	  // If the stack still contains elements, then it is an invalid expression.
// 	  return stack.isEmpty();
// 	}
// }
