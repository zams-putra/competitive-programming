package weeklycontest510

import (
	"strconv"
	"strings"
)

// Question :

// You are given two valid times startTime and endTime, each represented as a string in the format "HH:MM:SS".
// Return the number of seconds that have elapsed from startTime to endTime, inclusive of both endpoints.
//
// Example 1:
// Input: startTime = "01:00:00", endTime = "01:00:25"
// Output: 25
// Explanation:
// endTime is 25 seconds ahead of startTime.
// Example 2:
// Input: startTime = "12:34:56", endTime = "13:00:00"
// Output: 1504
// Explanation:
// endTime is 25 minutes and 4 seconds ahead of startTime, which equals 1504 seconds.
//
// Constraints:
// startTime.length == 8
// endTime.length == 8
// startTime and endTime are valid times in the format "HH:MM:SS"
// 00 <= HH <= 23
// 00 <= MM <= 59
// 00 <= SS <= 59
// endTime is not earlier than startTime

// Answer :

func NumOfElapsedSecsBetweenTwoTimes(startTime string, endTime string) int {
	arrX := strings.Split(startTime, ":")
	arrY := strings.Split(endTime, ":")
	x, y := 0, 0
	for i := 0; i < len(arrX); i++ {
		tX, _ := strconv.Atoi(arrX[i])
		tY, _ := strconv.Atoi(arrY[i])
		switch i {
		case 0:
			x += tX * 3600
			y += tY * 3600
		case 1:
			x += tX * 60
			y += tY * 60
		case 2:
			x += tX
			y += tY
		}
	}
	return y - x
}
