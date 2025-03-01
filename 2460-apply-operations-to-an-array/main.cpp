class Solution {
public:
    vector<int> applyOperations(vector<int>& nums) {
        
        size_t n = nums.size();

        for (size_t i = 0; i < n - 1; i++) {
            if (nums[i] == nums[i+1]) {
                nums[i] *= 2;
                nums[i + 1] = 0;
            }
        }

        size_t l = 0;
        for (size_t r = 0; r < n; r++) {
            if (nums[r] != 0) {
                swap(nums[r], nums[l]);
                l++;
            }
        }

        return nums;
    }
};
