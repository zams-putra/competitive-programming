#include <stdio.h>


// desc

// You are given two integers n and k.
// A positive integer x is called compatible if it satisfies both of the following conditions:
// abs(n - x) <= k
// (n & x) == 0
// Return the sum of all compatible integers x.
// Note:
// Here, & denotes the bitwise AND operator.
// The absolute difference between integers i and j is defined as abs(i - j).
//
// Example 1:
// Input: n = 2, k = 3
// Output: 10
// Explanation:
// The compatible integers are:
// x = 1, since abs(2 - 1) = 1 and 2 & 1 = 0.
// x = 4, since abs(2 - 4) = 2 and 2 & 4 = 0.
// x = 5, since abs(2 - 5) = 3 and 2 & 5 = 0.
// Thus, the answer is 1 + 4 + 5 = 10.
// Example 2:
// Input: n = 5, k = 1
// Output: 0
// Explanation:
// There are no compatible integers in the range [4, 6]. Thus, the answer is 0.
//
// Constraints:
// 1 <= n <= 100
// 1 <= k <= 100

// ans
int SumOfCompatibleNumsInRangeI(int n, int k) {
    int x = 1;
    int r = 0;
    if(n - k > x) {
        x = n - k;
    }
    for (int i = x; i <= n + k; i++){
        if(abs(n - i) <= k && (n & i) == 0) {
            r += i;
        }
    }
    return r;
}


int abs(int n) {
    if(n < 0) {
        return -n;
    }
    return n;
}