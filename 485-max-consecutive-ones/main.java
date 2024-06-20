class Solution {
    public int findMaxConsecutiveOnes(int[] nums) {
        
        int max_sum = 0;
        int curr_sum = 0;

        for (int i = 0; i < nums.length; i++) {
            if (nums[i] == 1) {
                curr_sum++;
                max_sum = Math.max(max_sum, curr_sum);
            }else {
                curr_sum = 0;
            }
        }

        return max_sum;
    }
}
