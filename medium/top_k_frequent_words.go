package medium

import (
	"container/heap"
	"strings"
)

// Given a non-empty list of words, return the k most frequent elements.
//
// Your answer should be sorted by frequency from highest to lowest. If
// two words have the same frequency, then the word with the lower
// alphabetical order comes first.
//
// Example 1:
// Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
// Output: ["i", "love"]
// Explanation: "i" and "love" are the two most frequent words.
//     Note that "i" comes before "love" due to a lower alphabetical order.
//
// Example 2:
// Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
// Output: ["the", "is", "sunny", "day"]
// Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
//     with the number of occurrence being 4, 3, 2 and 1 respectively.
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Input words contain only lowercase letters.
//
// Follow up:
// Try to solve it in O(n log k) time and O(n) extra space.

func topKFrequentWords(words []string, k int) []string {
	if len(words) == 0 || k <= 0 {
		return nil
	}

	var (
		i        int
		wordFreq = wordUniq(words)
		hp       = new(myMinHeap)
	)
	for word, freq := range wordFreq {
		if i < k {
			heap.Push(hp, &element{word: word, freq: freq})
			i++
		} else {
			if freq > (*hp)[0].freq ||
				(freq == (*hp)[0].freq && strings.Compare(word, (*hp)[0].word) == -1) {
				heap.Pop(hp)
				heap.Push(hp, &element{word: word, freq: freq})
			}
		}
	}

	var res []string
	for hp.Len() != 0 {
		res = append(res, heap.Pop(hp).(*element).word)
	}
	reverse(res)

	return res
}

// wordUniq returns a map which mapping word to frequent.
func wordUniq(words []string) map[string]int {
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word] += 1
	}
	return wordFreq
}

func reverse(s []string) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type element struct {
	word string
	freq int
}

type myMinHeap []*element

func (h myMinHeap) Len() int {
	return len(h)
}

func (h myMinHeap) Less(i, j int) bool {
	if h[i].freq < h[j].freq {
		return true
	}

	if h[i].freq == h[j].freq && strings.Compare(h[i].word, h[j].word) == 1 {
		return true
	}

	return false
}

func (h myMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *myMinHeap) Push(v interface{}) {
	*h = append(*h, v.(*element))
}

func (h *myMinHeap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
