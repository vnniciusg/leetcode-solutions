class Solution {
public:
    bool check(vector<int>& nums) {
        
        size_t n = nums.size();
        vector<int> sorted_nums(nums.begin(), nums.end());
        sort(sorted_nums.begin(), sorted_nums.end());

        for (size_t i = 0; i < n; i++){
            bool is_match = true;

            for (size_t j = 0; j < n; j++) {
                if (sorted_nums[j] != nums[(i + j) % n ]) {
                    is_match = false;
                    break;
                }
            }

            if (is_match) {
                return true;
            }
        }

        return false;
    }
};
