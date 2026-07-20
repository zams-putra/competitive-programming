package difficultyeasy

// Question :

// Given a 2D grid of size m x n and an integer k. You need to shift the grid k times.
// In one shift operation:
// Element at grid[i][j] moves to grid[i][j + 1].
// Element at grid[i][n - 1] moves to grid[i + 1][0].
// Element at grid[m - 1][n - 1] moves to grid[0][0].
// Return the 2D grid after applying shift operation k times.
// Example 1:
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 1
// Output: [[9,1,2],[3,4,5],[6,7,8]]
// Example 2:
// Input: grid = [[3,8,1,9],[19,7,2,5],[4,6,11,10],[12,0,21,13]], k = 4
// Output: [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
// Example 3:
// Input: grid = [[1,2,3],[4,5,6],[7,8,9]], k = 9
// Output: [[1,2,3],[4,5,6],[7,8,9]]
// Constraints:
// m == grid.length
// n == grid[i].length
// 1 <= m <= 50
// 1 <= n <= 50
// -1000 <= grid[i][j] <= 1000
// 0 <= k <= 100

// Answer :
func Shift2DGrid(grid [][]int, k int) (r [][]int) {
	a := []int{}
	for _, gr := range grid {
		for _, g := range gr {
			a = append(a, g)
		}
	}
	k %= len(a)
	newA, n := []int{}, len(a)-k
	for len(newA) < len(a) {
		if n == len(a) {
			n = 0
		}
		newA = append(newA, a[n])
		n++
	}
	temp := []int{}
	for i := 0; i < len(newA); i++ {
		temp = append(temp, newA[i])
		if len(temp) == len(grid[0]) {
			r = append(r, temp)
			temp = []int{}
		}
	}
	return
}
