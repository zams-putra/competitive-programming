// WPU Coding challenge
// 19/366

// Question :
// Given an array of integers as strings and numbers, return the sum of the array values as if all were numbers.
// Return your answer as a number.

// Answer :
package soal

import "strconv"

func SumMix(arr []any) int {
	result := 0
	for _, r := range arr {
		str, ok := r.(string)

		if ok{ // if string
			nowInt, _ := strconv.Atoi(str)
			result += nowInt
		} else {
			result += r.(int)
		}
	}

	return result
}

// func SumMix(arr []any) int {
// 	result := 0
// 	for _, r := range arr {
// 		switch data := r.(type) {
// 		case int:
// 			result += data
// 		case string:
// 			notStr,_ := strconv.Atoi(data)
// 			result += notStr
// 		}
// 	}

// 	return result
// }

// Done