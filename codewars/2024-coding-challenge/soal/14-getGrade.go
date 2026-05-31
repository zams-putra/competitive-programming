// WPU Coding challenge
// 14/366

// Question :
// Complete the function so that it finds the average of the three scores passed to it and returns the letter value associated with that grade.
// Numerical Score	Letter Grade
// 90 <= score <= 100	'A'
// 80 <= score < 90	'B'
// 70 <= score < 80	'C'
// 60 <= score < 70	'D'
// 0 <= score < 60	'F'
// Tested values are all between 0 and 100. Theres is no need to check for negative values or values greater than 100.

// func GetGrade(a,b,c int) rune {
// }

// rumus rata rata : temp += arr[i] / len(arr)

// Answer :
package soal

func GetGrade(a, b, c int) rune {
	total := a + b + c
	total /= 3
	var grade rune
	if total <= 100 && total >= 90{
		grade = 'A'
	} else if total < 90 && total >= 80 {
		grade = 'B'
	} else if total < 80 && total >= 70 {
		grade = 'C'
	} else if total < 70 && total >= 60 {
		grade = 'D'
	} else {
		grade = 'F'
	}
	// fmt.Println(string(grade))
	return grade
}

// Done

// atau :
// func GetGrade(a, b, c int) rune {
//     total := (a + b + c) / 3
//     var grade rune
//     switch {
//     case total >= 90 && total <= 100:
//         grade = 'A'
//     case total >= 80 && total < 90:
//         grade = 'B'
//     case total >= 70 && total < 80:
//         grade = 'C'
//     case total >= 60 && total < 70:
//         grade = 'D'
//     default:
//         grade = 'F'
//     }
//     return grade
// }
