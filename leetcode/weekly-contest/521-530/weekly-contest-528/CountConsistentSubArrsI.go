package weeklycontest528

import "slices"

// Question :

// You are given an integer array nums.
// A subarray is called consistent if the bitwise AND of all its elements is equal to the minimum value of the subarray.
// Return the number of consistent subarrays in nums​​​​​​​.
// A subarray is a contiguous, non-empty sequence of elements within an array.
//
// Example 1:
// Input: nums = [5,3,1,2]
// Output: 6
// Explanation:
// The following subarrays are consistent:
// Subarray	Values	Bitwise AND	Minimum Value
// nums[0..0]	[5]	5	5
// nums[1..1]	[3]	3	3
// nums[2..2]	[1]	1	1
// nums[3..3]	[2]	2	2
// nums[1..2]	[3, 1]	3 AND 1 = 1	1
// nums[0..2]	[5, 3, 1]	5 AND 3 AND 1 = 1	1
// Therefore, the number of consistent subarrays is 6.
// Example 2:
// Input: nums = [1,1,1]
// Output: 6
// Explanation:
// For every subarray, both the bitwise AND and the minimum value are equal to 1. Therefore, the answer is 6.
// Example 3:
// Input: nums = [0]
// Output: 1
// Explanation:
// The only subarray is [0], whose bitwise AND and minimum value are both 0. Therefore, the answer is 1.
//
// Constraints:
// 1 <= nums.length <= 100
// 0 <= nums[i] <= 100

// Answer :

func CountConsistentSubArrsI(nums []int) (r int) {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if j+i >= len(nums) {
				break
			}
			a := nums[j : (j+i)+1]
			bitwiseVal := a[0]
			for _, v := range a {
				bitwiseVal &= v
			}
			if bitwiseVal == slices.Min(a) {
				r++
			}
		}
	}
	return
}
