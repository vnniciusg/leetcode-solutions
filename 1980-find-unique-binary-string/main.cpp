class Solution {
public:
    string findDifferentBinaryString(vector<string>& nums) {
        vector<char> ans;
        size_t n = nums.size();

        for (size_t i = 0; i < n; i++) {
            char curr = nums[i][i];
            ans.push_back(curr == '0' ? '1' : '0');
        }

        return string(ans.begin(), ans.end());
    }
};
