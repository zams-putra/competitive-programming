//  WPU Coding challenge
// 2/366

// Quest :
// Build a function that returns an array of integers from n to 1 where n>0.
// Example : n=5 --> [5,4,3,2,1]

// func ReverseSeq(n int) []int {
//   return make([]int, n)
// }

// Answer :
package soal

func ReverseSeq(n int) []int {
	var kotak = []int{}
	for i := 0; i < n; n--{
	  kotak = append(kotak, n)
	}	
	return kotak
 }

//   Done