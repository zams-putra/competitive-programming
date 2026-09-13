package weeklycontest519

// Question :

// You are given an integer n, a 2D integer array grid of size n x n, and two integer arrays rowShift and colShift, each of length n where:
// rowShift[i] represents the number of positions to cyclically left shift the ith row of grid.
// colShift[j] represents the number of positions to cyclically upward shift the jth column of grid.
// First, cyclically shift each row according to rowShift, then cyclically shift each column according to colShift.
// Return the resulting grid after performing all the shifts.
// A cyclic left shift of the ith row by k positions shifts only that row. The element at column j moves to column (j - k + n) % n, while all other rows remain unchanged.
// A cyclic upward shift of the jth column by k positions shifts only that column. The element at row i moves to row (i - k + n) % n, while all other columns remain unchanged.
//
// Example 1:
// Input: n = 2, grid = [[1,2],[3,4]], rowShift = [1,0], colShift = [0,1]
// Output: [[2,4],[3,1]]
// Explanation:
// The grid changes as follows:
// Example 2:
// Input: n = 3, grid = [[1,2,3],[4,5,6],[7,8,9]], rowShift = [1,2,0], colShift = [2,2,1]
// Output: [[7,8,5],[2,3,9],[6,4,1]]
// Explanation:
// The grid changes as follows:
//
// Constraints:
// 1 <= n == grid.length == grid[i].length <= 10
// 1 <= grid[i][j] <= 100
// rowShift.length == colShift.length == n
// 0 <= rowShift[i], colShift[i] < n

// Answer :

func CyclicallyShiftRowsAndCols(n int, grid [][]int, rowShift []int, colShift []int) (result [][]int) {
	temp := [][]int{}
	for i := 0; i < len(grid); i++ {
		newRow := []int{}
		for j := rowShift[i]; j < len(grid[i]); j++ {
			newRow = append(newRow, grid[i][j])
		}
		for j := 0; j < rowShift[i]; j++ {
			newRow = append(newRow, grid[i][j])
		}
		temp = append(temp, newRow)

	}
	result = make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newRow := (i - colShift[j] + n) % n
			result[newRow][j] = temp[i][j]
		}
	}
	return
}
