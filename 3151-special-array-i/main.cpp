class Solution {
public:
    bool isArraySpecial(vector<int>& nums) {
        
        for (size_t i = 1; i < nums.size(); i++) {
            if (((nums[i] & 1) ^ (nums[i - 1] & 1)) == 0) {
                return false;
            }
        }

        return true;
    }
};
