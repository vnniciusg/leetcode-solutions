class Solution {
public:
    int minMoves(vector<int>& nums) {
        int sum = std::accumulate(nums.begin(), nums.end(), 0);
        int min_value = *std::min_element(nums.begin(), nums.end());
        return sum - (nums.size() * min_value);
    }
};