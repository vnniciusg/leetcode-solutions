class Solution {
public:
    int xorAllNums(vector<int>& nums1, vector<int>& nums2) {
        int xor1 = 0, xor2 = 0;
        size_t len1 = nums1.size(), len2 = nums2.size();

        if (len2 & 1 != 0) {
            for (const auto& num: nums1) {
                xor1 ^= num;
            }
        }

        if (len1 & 1 != 0) {
            for (const auto& num: nums2) {
                xor2 ^= num;
            }
        }

        return xor1 ^ xor2;
    }
};
