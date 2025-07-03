class Solution {
public:
    int missingNumber(vector<int>& nums) {
        int result = 0;

        for (int i = 0; i < nums.size() + 1; i++) {
            result ^= i;
        }

        for (const int num: nums) {
            result ^= num;
        }

        return result;
    }
};
