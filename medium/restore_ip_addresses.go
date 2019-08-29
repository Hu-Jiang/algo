package medium

import "strconv"

// Given a string containing only digits, restore it by returning
// all possible valid IP address combinations.
//
// Example:
// Input: "25525511135"
// Output: ["255.255.11.135", "255.255.111.35"]

func restoreIpAddresses(s string) []string {
	addr := new(IpAddresses)
	addr.backtrack("", s, 0)
	return *addr
}

type IpAddresses []string

func (a *IpAddresses) backtrack(prev, left string, depth int) {
	if depth == 4 {
		if len(left) == 0 {
			*a = append(*a, prev)
		}
		return
	}

	if len(prev) != 0 {
		prev += "."
	}
	for i := 1; i < 4; i++ {
		if len(left) < i {
			break
		}
		if i > 1 && left[0] == '0' { // skip leader 0
			break
		}

		value, _ := strconv.Atoi(left[:i])
		if value < 256 {
			a.backtrack(prev+left[:i], left[i:], depth+1)
		}
	}
}
