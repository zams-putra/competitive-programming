// WPU Coding challenge
// 15/366

// Question :
// Task
// Sum all the numbers of a given array ( cq. list ), except the highest and the lowest element ( by value, not by index! ).
// The highest or lowest element respectively is a single element at each edge, even if there are more than one with the same value.
// Mind the input validation.
// Example
// { 6, 2, 1, 8, 10 } => 16
// { 1, 1, 11, 2, 3 } => 6
// Input validation
// If an empty value ( null, None, Nothing etc. ) is given instead of an array, or the given array is an empty list or a list with only 1 element, return 0.

// Answer :
package soal

func SumArray(arr []int) int {
	if arr == nil || len(arr) <= 2 {
		return 0
	}
	temp := []int{}

	for i := 0; i < len(arr) - 1; i++{
		for j := 0; j < len(arr) - i - 1; j++{
			if arr[j] > arr[j + 1] {
				arr[j], arr[j + 1] = arr[j + 1], arr[j]
			}
		}
	}

	for i := 1; i < len(arr) - 1; i++{
		temp = append(temp, arr[i])
	}

	var result int

	for _, r := range temp {
		result += r
	}

	return result

}

// Done

// versi js :
// function sumArray(array) {
// 	if (!Array.isArray(array)|| array.length <= 2) return 0;

// 	let temp = []

// 	for (let i = 0; i < array.length - 1; i++) {
// 	  for (let j = 0; j < array.length - i - 1; j++){
// 		if(array[j] > array[j + 1]){
// 		   let tempval = array[j]
// 		   array[j] = array[j + 1]
// 		  array[j + 1] = tempval
// 		}
// 	  }
// 	}

// 	for (let i = 1; i < array.length - 1; i++){
// 	  temp.push(array[i])
// 	}

// 	let result = 0

// 	for (let i = 0; i < temp.length; i++){
// 	  result += temp[i]
// 	}

// 	return result
//   }