package weeklycontest516

import "fmt"

// Question :

// You are given a string s consisting of lowercase English letters.
// Construct a binary string by replacing each character in s with the 8-bit binary representation of its ASCII value, including leading zeros, while preserving the original order of the characters.
// Return true if the resulting binary string is a palindrome. Otherwise, return false.
// A binary string is a string which contains only the characters '0' and '1'.
// A palindrome is a string that reads the same forward and backward.
//
// Example 1:
// Input: s = "ff"
// Output: true
// Explanation:
// The ASCII value of f is 102, whose 8-bit binary representation is 01100110.
// Thus, the binary string is 0110011001100110.
// Since this binary string is a palindrome, the output is true.
// Example 2:
// Input: s = "leet"
// Output: false
// Explanation:
// The ASCII values of l, e, e, and t are 108, 101, 101, and 116, respectively.
// Their 8-bit binary representations are 01101100, 01100101, 01100101, and 01110100.
// Thus, the binary string is 01101100011001010110010101110100.
// Since this binary string is not a palindrome, the output is false.
//
// Constraints:
// 1 <= s.length <= 100
// s consists of lowercase English letters.

// Answer :
// instead of pake a := strconv.FormatInt(102, 2)
// disini aku pake  a := fmt.Sprintf("%08b", 102), biar bisa 8 bit soalnya kalai yg atas dia ga include 0 di depan misal 0110 jadi 110
func CheckASCIIPalindromic(s string) bool {
	bin := ""
	for i := 0; i < len(s); i++ {
		temp := fmt.Sprintf("%08b", int(s[i]))
		bin += temp
	}
	for i := 0; i < len(bin)/2; i++ {
		x, y := string(bin[i]), string(bin[len(bin)-(i+1)])
		if x != y {
			return false
		}
	}
	return true
}
