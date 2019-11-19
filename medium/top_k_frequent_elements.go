package medium

import "container/heap"

// Given a non-empty array of integers, return the k most frequent elements.
//
// Example 1:
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
//
// Example 2:
// Input: nums = [1], k = 1
// Output: [1]
//
// Note:
// You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
// Your algorithm's time complexity must be better than O(n log n), where
// n is the array's size.

func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 || k <= 0 {
		return nil
	}

	var (
		i       int
		numFreq = uniq(nums)
		hp      = new(minHeap)
	)
	for val, freq := range numFreq {
		if i < k {
			heap.Push(hp, &item{val: val, freq: freq})
			i++
		} else if freq > (*hp)[0].freq {
			heap.Pop(hp)
			heap.Push(hp, &item{val: val, freq: freq})
		}
	}

	var res []int
	for hp.Len() != 0 {
		res = append(res, heap.Pop(hp).(*item).val)
	}
	reverseSlice(res)

	return res
}

// uniq returns a map which mapping number to frequent.
func uniq(nums []int) map[int]int {
	numFreq := make(map[int]int)
	for _, v := range nums {
		numFreq[v] += 1
	}
	return numFreq
}

func reverseSlice(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

type item struct {
	val  int
	freq int
}

type minHeap []*item

func (h minHeap) Len() int {
	return len(h)
}

func (h minHeap) Less(i, j int) bool {
	return h[i].freq < h[j].freq
}

func (h minHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *minHeap) Push(v interface{}) {
	*h = append(*h, v.(*item))
}

func (h *minHeap) Pop() interface{} {
	v := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return v
}
