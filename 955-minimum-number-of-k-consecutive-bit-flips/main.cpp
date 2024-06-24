class Solution {
public:
    int minKBitFlips(vector<int>& nums, int k) {
        
        int n = nums.size();
        int flips = 0;
        int flip_effect = 0;
        std::vector<int> is_flipped(n, 0);

        for (int i = 0; i < n; ++i) {
            if (i >= k) {
                flip_effect ^= is_flipped[i - k];
            }

            if (flip_effect == nums[i]) {
                if (i + k > n) {
                    return -1;
                }
                flips++;
                flip_effect ^= 1;
                is_flipped[i] = 1;
            }
        }

        return flips;

    }   
};
