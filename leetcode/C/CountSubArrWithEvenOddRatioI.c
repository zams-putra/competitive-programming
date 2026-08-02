int CountSubArrWithEvenOddRatioI(int* nums, int numsSize, int a, int b) {
    int r = 0;
    for(int i = 0; i < numsSize; i++){
        int even = 0;
        int odd = 0;
        for(int j = i; j < numsSize; j++){
            if(nums[j] % 2 == 0) {
                even++;
            } else {
                odd++;
            }
            if(odd > 0 && even*b <= a*odd){
                r++;
            }
        }
    }
    return r;
}