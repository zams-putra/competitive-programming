// WPU Coding challenge
// 17/366

// Question :
// You are given the length and width of a 4-sided polygon. The polygon can either be a rectangle or a square.
// If it is a square, return its area. If it is a rectangle, return its perimeter.
// Example(Input1, Input2 --> Output):
// 6, 10 --> 32
// 3, 3 --> 9
// Note: for the purposes of this kata you will assume that it is a square if its length and width are equal, otherwise it is a rectangle.

// const areaOrPerimeter = function(l , w) {
// };

// Answer :
package soal

func AreaOrPerimeter(w, l int) int {
	if w == l {
		return w * l
	} else {
		return w * 2 + l * 2
	}
}

// Done

// soal aslinya pake js
// versi js :
// const areaOrPerimeter = (l , w) => l === w ? l * w : l * 2 + w * 2
