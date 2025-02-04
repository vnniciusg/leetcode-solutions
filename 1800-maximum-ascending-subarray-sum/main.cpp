class Solution {
public:
    int maxAscendingSum(vector<int>& nums) {
        
        int _max = 0;
        int currSubSum = nums[0];

        for (size_t idx = 1; idx < nums.size(); idx++) {
            
            if (nums[idx] <= nums[idx - 1]) {
                _max = std::max(_max, currSubSum);
                currSubSum = 0;
            }

            currSubSum += nums[idx];
        }

        return std::max(_max, currSubSum);
    }
};
