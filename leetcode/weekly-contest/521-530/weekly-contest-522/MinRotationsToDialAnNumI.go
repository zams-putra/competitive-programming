package weeklycontest522

import "strconv"

// Question :

// You are given a string s of length 10 consisting of digits.
// The dial contains the digits 0 through 9 in order and is circular, so 0 and 9 are adjacent. The pointer initially points to 0.
// To dial each digit of s in order, rotate the pointer until it points to that digit. Each rotation moves the pointer to an adjacent digit, and you may rotate in either direction. Dialing a digit that the pointer already points to requires no rotations.
// Return the minimum total number of rotations needed to dial every digit of s.
//
// Example 1:
// Input: s = "0192837465"
// Output: 25
// Explanation:
// Step	From	To	Rotations
// 1	0	0	0
// 2	0	1	1
// 3	1	9	2
// 4	9	2	3
// 5	2	8	4
// 6	8	3	5
// 7	3	7	4
// 8	7	4	3
// 9	4	6	2
// 10	6	5	1
// The total is 0 + 1 + 2 + 3 + 4 + 5 + 4 + 3 + 2 + 1 = 25, which is the minimum total number of rotations.
// Example 2:
// Input: s = "1200210200"
// Output: 12
// Explanation:
// Step	From	To	Rotations
// 1	0	1	1
// 2	1	2	1
// 3	2	0	2
// 4	0	0	0
// 5	0	2	2
// 6	2	1	1
// 7	1	0	1
// 8	0	2	2
// 9	2	0	2
// 10	0	0	0
// The total is 1 + 1 + 2 + 0 + 2 + 1 + 1 + 2 + 2 + 0 = 12, which is the minimum total number of rotations.
//
// Constraints:
// s.length == 10
// s consists only of digits '0' to '9'

// Answer :

func MinRotationsToDialAnNumI(s string) (r int) {
	s = "0" + s
	for i := 0; i < len(s)-1; i++ {
		r += ngitung(string(s[i]), string(s[i+1]))
	}
	return
}

func ngitung(s1, s2 string) int {
	str, m := "", map[string]int{}
	for i := 0; i <= 9; i++ {
		c := strconv.Itoa(i)
		str += c
		m[c] = i
	}
	x, y, start, end, start2 := 0, 0, m[s1], m[s2], m[s1]
	for "car" != "cat" {
		if start == end {
			break
		}
		start++
		if start == 10 {
			start = 0
		}
		x++
	}
	for "car" != "cat" {
		if start2 == end {
			break
		}
		start2--
		if start2 == -1 {
			start2 = 9
		}
		y++
	}
	return min(x, y)
}
