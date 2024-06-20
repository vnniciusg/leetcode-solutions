public class Solution {
    public int FindMaxConsecutiveOnes(int[] nums) {
        
        int max_sum = 0;
        int curr_sum = 0;

        foreach(var num in nums) 
        {
            if (num == 1)
            {
                curr_sum++;
                max_sum = Math.Max(max_sum, curr_sum);
            }

            else
            {
                curr_sum = 0;
            }

        }

        return max_sum;
    }

}
