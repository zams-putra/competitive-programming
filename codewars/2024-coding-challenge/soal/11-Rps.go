// WPU Coding challenge
// 11/366

// Question :
// Rock Paper Scissors
// Let's play! You have to return which player won! In case of a draw return Draw!.

// Examples(Input1, Input2 --> Output):

// "scissors", "paper" --> "Player 1 won!"
// "scissors", "rock" --> "Player 2 won!"
// "paper", "paper" --> "Draw!"

// func Rps(p1, p2 string) string {
// 	return "Draw!"
//   }

// Answer :
package soal

func Rps(p1, p2 string) string {
	var result string
	if p1 == p2 {
		result = "Draw!"
	} else if p1 == "scissors" {
		if p2 == "rock" {
			result = "Player 2 won!"
		} else {
			result = "Player 1 won!"
		}
	} else if p1 == "rock" {
		if p2 == "paper" {
			result = "Player 2 won!"
		} else {
			result = "Player 1 won!"
		}
	} else if p1 == "paper" {
		if p2 == "scissors" {
			result = "Player 2 won!"
		} else {
			result = "Player 1 won!"
		}
	}
	return result
}

// Done

// func Rps(p1, p2 string) string {
// 	if p1 == p2 {
// 		return "Draw!"
// 	} else if p1 == "rock" && p2 == "scissors"{
// 		return "Player 1 won!"
// 	} else if p1 == "scissors" && p2 == "paper"{
// 		return "Player 1 won!"
// 	} else if p1 == "paper" && p2 == "rock" {
// 		return "Player 1 won!"
// 	} else {
// 		return "Player 2 won!"
// 	}
// }

// func Rps(p1, p2 string) string {
// 	if p1 == p2 {
// 		return "Draw!"
// 	} 
// 	if p1 == "rock" && p2 == "scissors" || p1 == "scissors" && p2 == "paper" || p1 == "paper" && p2 == "rock"{
// 		return "Player 1 won!"
// 	} else {
// 		return "Player 2 won!"
// 	}
// }