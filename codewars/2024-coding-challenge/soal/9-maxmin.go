// WPU Coding challenge
// 9/366

// Question :

// Your task is to make two functions ( max and min, or maximum and minimum, etc., depending on the language ) that receive a list of integers as input, and return the largest and lowest number in that list, respectively.

// Examples (Input -> Output)
// * [4,6,2,1,9,63,-134,566]         -> max = 566, min = -134
// * [-52, 56, 30, 29, -54, 0, -110] -> min = -110, max = 56
// * [42, 54, 65, 87, 0]             -> min = 0, max = 87
// * [5]                             -> min = 5, max = 5

// var min = function(list){
//     return list[0];
// }

// var max = function(list){
//     return list[0];
// }

// Answer :
package soal

// func Max(arr []int) int {
// 	var result int

// 	for i := 0; i < len(arr); i++{
// 		temp := len(arr) - 1
// 		if arr[temp] > arr[i] {
// 			result = arr[temp]
// 		}
// 	}
// 	return result
// }

// func Min(arr []int) int {
// 	left := 0
// 	var result int
// 	for i := 1; i < len(arr); i++{
// 		right := i

// 		if arr[right] < arr[left]{
// 			result = arr[right]
// 		}
// 	}
// 	return result
// }

// Ternyata gini : revisi :
func Max(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	var result int
	left := 0

	for i := 1; i < len(arr); i++{
		if arr[i] > arr[left] {
			result = arr[i]
			left = i
		}
	}
	return result
}

func Min(arr []int) int {

	if len(arr) == 1 {
		return arr[0]
	}

	left := 0
	var result int
	for i := 1; i < len(arr); i++{
		right := i

		if arr[right] < arr[left]{
			result = arr[right]
		}
	}
	return result
}

// Done
// aslinya gada go di soalnya
// versi js :
// var min = function(list){
// 	let minVal = list[0]
// 	for(let i = 1; i < list.length; i++){
// 	  if(minVal > list[i]){
// 		minVal = list[i]
// 	  }
// 	}
// 	return minVal
//   }

//   var max = function(list){
// 	let maxVal = list[0]
// 	for(let i = 1; i < list.length; i++){
// 	  if(maxVal < list[i]){
// 		maxVal = list[i]
// 	  }
// 	}
// 	return maxVal
//   }


// Versi c++ :
// int min(vector<int> list){
//     int left = list[0];
//     for (int i = 1; i < list.size(); i++){
//       int right = list[i];
//       if (right < left){
//         left = right;
//       }
//     }
//     return left;
// }

// int max(vector<int> list){
//     int left = list[0];
//     for (int i = 1; i < list.size(); i++){
//       int right = list[i];
//       if (right > left){
//         left = right;
//       }
//     }
//     return left;
// }