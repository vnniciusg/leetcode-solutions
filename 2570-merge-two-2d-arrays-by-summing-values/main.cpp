class Solution {
public:
    vector<vector<int>> mergeArrays(vector<vector<int>>& nums1, vector<vector<int>>& nums2) {
        
        unordered_map<int, int> _dict;

        for (const auto& nums: nums1) {
            _dict[nums[0]] = nums[1];
        }

        for (const auto& nums: nums2) {
            _dict[nums[0]] += nums[1];
        }

        vector<vector<int>> ans;
        for (const auto& [key, value]: _dict) {
            ans.push_back({key, value});
        }

        sort(ans.begin(), ans.end());
        return ans;
    }
};
