// WPU Coding challenge
// 1/366

//  Quest :
// Given a non-empty array of integers, return the result of multiplying the values together in order. Example:
// [1, 2, 3, 4] => 1 * 2 * 3 * 4 = 24

// func Grow(arr []int) int{

// }

// Answer :
package soal

func Grow(arr []int) int{
	var final int = 1

	for i := 0; i < len(arr); i++{
		final *= arr[i]
	}
	
	return final
}

//            Done