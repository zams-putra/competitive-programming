package difficultyeasy

// Question :

// You are given a 0-indexed array of integers nums.
// A prefix nums[0..i] is sequential if, for all 1 <= j <= i, nums[j] = nums[j - 1] + 1. In particular, the prefix consisting only of nums[0] is sequential.
// Return the smallest integer x missing from nums such that x is greater than or equal to the sum of the longest sequential prefix.
// Example 1:
// Input: nums = [1,2,3,2,5]
// Output: 6
// Explanation: The longest sequential prefix of nums is [1,2,3] with a sum of 6. 6 is not in the array, therefore 6 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
// Example 2:
// Input: nums = [3,4,5,1,12,14,13]
// Output: 15
// Explanation: The longest sequential prefix of nums is [3,4,5] with a sum of 12. 12, 13, and 14 belong to the array while 15 does not. Therefore 15 is the smallest missing integer greater than or equal to the sum of the longest sequential prefix.
// Constraints:
// 1 <= nums.length <= 50
// 1 <= nums[i] <= 50

// Answer :

func SmallestMissingIntGreaterThanSeqPrefixSum(nums []int) int {
	arr := []int{}
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) {
			if nums[i+1]-nums[i] == 1 {
				arr = append(arr, nums[i])
			} else {
				if i != 0 && nums[i]-nums[i-1] == 1 {
					arr = append(arr, nums[i])
				}
				break
			}
		} else {
			if i == 0 {
				arr = append(arr, nums[i])
				break
			}
			if nums[i]-nums[i-1] == 1 {
				arr = append(arr, nums[i])
			}
		}
	}
	if len(arr) == 0 {
		arr = append(arr, nums[0])
	}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] == sum {
				sum = nums[j] + 1
			}
		}
	}
	return sum
}
