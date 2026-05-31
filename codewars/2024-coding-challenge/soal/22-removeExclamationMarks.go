// WPU Coding challenge
// 22/366

// Question :

// Write function RemoveExclamationMarks which removes all exclamation marks from a given string.

// Answer :
package soal

func RemoveExclamationMarks(text string) (res string) {

	for _, t := range text {
		if string(t) != "!" {
			res += string(t)
		}
	}

	return res

}


// Aslinya pakai js gada go : 
// function removeExclamationMarks(s) {
// 	let res = ""
	
// 	for (let i = 0; i < s.length; i++){
// 	  if (s[i] != "!") {
// 		res += s[i]
// 	  }
// 	}
// 	return res
//   }

// let removeExclamationMarks = (s) => s.split("!").join("")