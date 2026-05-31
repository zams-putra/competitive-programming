package soal

import "strconv"

// desc
// Write a function which takes a number as input and returns the sum of the absolute value of each of the number's decimal digits.
// For example: (Input --> Output)
// 10 --> 1
// 99 --> 18
// -32 --> 5
// Let's assume that all numbers in the input will be integer values.

// ans
func SummingAnNumbersDigits(number int) (r int) {
	s := strconv.Itoa(number)
	for _, v := range s {
		n, _ := strconv.Atoi(string(v))
		r += n
	}
	return 
}