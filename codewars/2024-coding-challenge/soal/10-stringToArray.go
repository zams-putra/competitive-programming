// WPU Coding challenge
// 10/366

// Question :

// Write a function to split a string and convert it into an array of words.

// Examples (Input ==> Output):
// "Robin Singh" ==> ["Robin", "Singh"]

// "I love arrays they are my favorite" ==> ["I", "love", "arrays", "they", "are", "my", "favorite"]

// func StringToArray(str string) []string {
// }

// Answer :
package soal

func StringToArray(str string) []string {
	temp := ""
	var arr []string
	for i := 0; i < len(str); i++{
		if str[i] == ' '{
			arr = append(arr, temp)
			temp = ""
		} else{
			temp += string(str[i])
		}
	}
	arr = append(arr, temp)
	return arr
}
// Done