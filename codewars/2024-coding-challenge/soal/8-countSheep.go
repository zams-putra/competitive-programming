// WPU Coding challenge
// 8/366

// Question :

// If you can't sleep, just count sheep!!

// Task:
// Given a non-negative integer, 3 for example, return a string with a murmur: "1 sheep...2 sheep...3 sheep...". Input will always be valid, i.e. no negative integers.

// func countSheep(num int) string {
// }

// Answer :
package soal

import (
	"strconv"
)

func CountSheep(num int) string {
  var result string
	for i := 1; i <= num; i++ {
		temp :=  strconv.Itoa(i)
		result += temp + " sheep..."
	}
	return result
}

// func countSheep(num int) string {
// 	var result string
// 	for i := 1; i <= num; i++ {
// 		result += strconv.Itoa(i) + " sheep..."
// 	}
// 	return result
// }

//  Done