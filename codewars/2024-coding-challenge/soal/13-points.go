// WPU Coding challenge
// 13/366

// Question :
// Our football team has finished the championship.
// Our team's match results are recorded in a collection of strings. Each match is represented by a string in the format "x:y", where x is our team's score and y is our opponents score.
// For example: ["3:1", "2:2", "0:1", ...]
// Points are awarded for each match as follows:
// if x > y: 3 points (win)
// if x < y: 0 points (loss)
// if x = y: 1 point (tie)
// We need to write a function that takes this collection and returns the number of points our team (x) got in the championship by the rules given above.
// Notes:
// our team always plays 10 matches in the championship
// 0 <= x <= 4
// 0 <= y <= 4

// func Points(games []string) int {
// 	return -1
//   }

// Answer :
package soal

func Points(games []string) int {
	result := 0
	for _, game := range games {
		x := game[0]
		y := game[2]
		if x > y {
			result += 3
		} else if x == y {
			result += 1
		} else {
			result += 0
		}
	}
	return result
}

// Done

// Atau :
// func Points(games []string) int {
// 	result := 0
// 	   for _, game := range games {
// 		   if game[0] > game[2] {
// 			   result += 3
// 		   } else if game[0] == game[2] {
// 			   result += 1
// 		   } else {
// 			   result += 0
// 		   }
// 	   }
// 	   return result
//    }

// func Points(games []string) int {
// 	result := 0
// 	for i := 0; i < len(games); i++ {
// 		if games[i][0] > games[i][2] {
// 			result += 3
// 		} else if games[i][0] == games[i][2] {
// 			result += 1
// 		} else {
// 			result += 0
// 		}
// 	}
// 	return result
// }