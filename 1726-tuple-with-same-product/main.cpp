class Solution {
public:
    int tupleSameProduct(vector<int>& nums) {
        
        size_t n = nums.size();
        unordered_map<int, int> freq;
        for (size_t i = 0; i < n - 1; i++) {
            for (size_t j = i + 1; j < n; j++) {
                freq[nums[i] * nums[j]]++;
            }
        }

        int count = 0;
        for (auto [product, frequency]: freq) {
            count += 8 * (frequency * (frequency - 1)) / 2;
        }

        return count;
    }
};
