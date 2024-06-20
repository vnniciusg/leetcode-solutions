/**
 * @param {number[]} nums
 * @return {number}
 */
var findMaxConsecutiveOnes = function(nums) {
    
    let max_sum = 0;
    let curr_sum = 0;

    for (i = 0; i < nums.length; i++) {
        if (nums[i] == 1 ){
            curr_sum++;
            max_sum = Math.max(max_sum, curr_sum);
        }else {
            curr_sum = 0;
        }
    }

    return max_sum;
};
