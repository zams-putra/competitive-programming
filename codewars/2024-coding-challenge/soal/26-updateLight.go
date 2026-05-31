// WPU Coding challenge
// 26/366

// Question :

// You're writing code to control your town's traffic lights. You need a function to handle each change from green, to yellow, to red, and then to green again.
// Complete the function that takes a string as an argument representing the current state of the light and returns a string representing the state the light should change to.
// For example, when the input is green, output should be yellow.

// Answer :
package soal

func UpdateLight(current string) string {
	switch {
	case current == "red":
		return "green"
	case current == "yellow":
		return "red"
	case current == "green":
		return "yellow"
	}

	return "Invalid"
}

// aslinya pake js soalnya gada go :
// const updateLight = (current) => current === "red" ? "green" : current === "yellow" ? "red" : "yellow"