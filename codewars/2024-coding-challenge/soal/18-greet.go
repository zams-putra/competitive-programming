// WPU Coding challenge
// 18/366

// Question :
// Create a function that gives a personalized greeting. This function takes two parameters: name and owner.
// Use conditionals to return the proper message:
// case	return
// name equals owner	'Hello boss'
// otherwise	'Hello guest'

// Answer :
package soal

func Greet(name, owner string) string {
	if name == owner {
		return "Hello boss"
	} else {
		return "Hello guest"
	}
}

// Aslinya pake js , versi js :
// let greet = (name, owner) => name === owner ? "Hello boss" : "Hello guest"