class Solution {
public:
    int findLHS(vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        int j = 0, maxLength = 0;

        for (int i = 0; i < nums.size(); i++) {
            while (nums[i] - nums[j] > 1) {
                j++;
            }

            if (nums[i] - nums[j] == 1) {
                maxLength = std::max(maxLength, i - j + 1);
            }
        }

        return maxLength;
    }
};
