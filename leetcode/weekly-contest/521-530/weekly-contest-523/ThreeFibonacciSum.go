package weeklycontest523

// Question :

// You are given an integer n.
// The Fibonacci sequence starts with 0 and 1, and each subsequent number is the sum of the previous two numbers. Its first few terms are 0, 1, 1, 2, 3, 5, 8, 13, ....
// Return true if n can be expressed as the sum of three consecutive terms in the Fibonacci sequence, and false otherwise.
//
// Example 1:
// Input: n = 16
// Output: true
// Explanation:
// The three consecutive terms 3, 5, and 8 have a sum of 3 + 5 + 8 = 16.
// Example 2:
// Input: n = 8
// Output: false
// Explanation:
// Although 1 + 2 + 5 = 8, these terms are not consecutive in the Fibonacci sequence. No three consecutive terms have a sum of 8.
// Example 3:
// Input: n = 2
// Output: true
// Explanation:
// The first three terms have a sum of 0 + 1 + 1 = 2.
//
// Constraints:
// 1 <= n <= 109

// Answer :

func ThreeFibonacciSum(n int) bool {
	arr := []int{}
	a, b := 0, 1
	for i := 0; i <= n; i++ {
		arr = append(arr, a)
		a, b = b, a+b
		if a > n {
			break
		}
	}
	for i := 0; i < len(arr); i++ {
		if i+2 < len(arr) {
			if arr[i]+arr[i+1]+arr[i+2] == n {
				return true
			}
		}
	}
	return false
}
