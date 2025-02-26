class Solution {
public:
    int maxAbsoluteSum(vector<int>& nums) {
        int maxSum = 0, minSum = 0, maxEnding = 0, minEnding = 0;

        for (const auto& num: nums) {
            maxEnding = std::max(maxEnding + num, num);
            minEnding = std::min(minEnding + num, num);

            maxSum = std::max(maxSum, maxEnding);
            minSum = std::min(minSum, minEnding);
        }

        return std::max(maxSum, std::abs(minSum));
    }
};
