// WPU Coding challenge
// 4/366

// Question :
// Write a function that takes an array of words and smashes them together into a sentence and returns the sentence. You can ignore any need to sanitize words or add punctuation, but you should add spaces between each word. Be careful, there shouldn't be a space at the beginning or the end of the sentence!

// ['hello', 'world', 'this', 'is', 'great']  =>  'hello world this is great'

// function smash (words) {
//     return ""
//  };

// Answer :
package soal
func Smash(words []string) string {
	var text = "'"
	for i := 0; i < len(words); i++ {
		text += words[i]
		if i != len(words)-1 {
			text += " "
		}
	}
	text += "'"
	return text
}


//  Done


// *note, soal asli pakai javascript, aku paksa pakai go : 
// Main : 

// arr := []string{"Hello", "World", "I", "Am", "Root"}
// 	fmt.Println(Smash(arr))


// ini js :
// function smash (words) {
  
// 	let text = ''
// 	 for (let i = 0; i < words.length; i++){
// 	   text += words[i]
// 	   if(i != words.length - 1) {
// 		 text += ' '
// 	   }
// 	 }
// 	return text
//   };