package usingimportpackage

import "strconv"

// Question :

// You are given an integer n.
// Return true if its binary representation contains exactly one pair of consecutive set bits, and false otherwise.
// Example 1:
// Input: nums = 6
// Output: true
// Explanation:
// Binary representation of 6 is 110.
// There is exactly one pair of consecutive set bits ("11"). Thus, the answer is true​​​​​​​.
// Example 2:
// Input: nums = 5
// Output: false
// Explanation:
// Binary representation of 5 is 101.
// There are no consecutive set bits. Thus, the answer is false​​​​​​​.
// Constraints:
// 0 <= n <= 105

// Answer :

func ExactlyOneConsecutiveSetBitsPair(n int) bool {
	s, c := strconv.FormatInt(int64(n), 2), 0
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] && s[i] == '1' {
			c++
		}
	}
	return c == 1
}
