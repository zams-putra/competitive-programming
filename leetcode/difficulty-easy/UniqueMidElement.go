package difficultyeasy

// Question :

// You are given an integer array nums of odd length n.
// Return true if the middle element of nums appears exactly once in the array. Otherwise return false.
// Example 1:
// Input: nums = [1,2,3]
// Output: true
// Explanation:
// The middle element of nums is 2, which appears exactly once.
// Thus, the answer is true.
// Example 2:
// Input: nums = [1,2,2]
// Output: false
// Explanation:
// The middle element of nums is 2, which appears twice.
// Thus, the answer is false.
// Constraints:
// 1 <= n == nums.length <= 100
// n is odd.
// 1 <= nums[i] <= 100

// Answer :

func UniqueMidElement(nums []int) bool {
	mid, c := nums[len(nums)/2], 0
	for _, v := range nums {
		if v == mid {
			c++
		}
	}
	return c == 1
}
