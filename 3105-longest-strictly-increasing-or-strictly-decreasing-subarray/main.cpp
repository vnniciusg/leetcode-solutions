class Solution {
public:
    int longestMonotonicSubarray(vector<int>& nums) {
        int inc_length = 1, dec_length = 1, max_length = 1;

        for (size_t i = 0; i < nums.size() - 1; i++) {
            if (nums[i + 1] > nums[i]) {
                inc_length++;
                dec_length = 1;

            }else if(nums[i + 1] < nums[i]) {
                dec_length++;
                inc_length = 1;
            }else {
                inc_length = dec_length = 1;
            }

            max_length = std::max(max_length, std::max(inc_length, dec_length));
        }

        return max_length;
    }
};
