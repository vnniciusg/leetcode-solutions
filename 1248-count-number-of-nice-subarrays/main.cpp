class Solution {
public:
    int numberOfSubarrays(vector<int>& nums, int k) {
        
        std::unordered_map<int, int> map;
        map[0] = 1;

        int count = 0;
        int odd_count = 0;

        for (const auto& num : nums) {
            if ((num & 1) != 0) {
                odd_count++;
            }

            if (map.find(odd_count - k) != map.end()) {
                count += map[odd_count - k];
            }

            map[odd_count]++;
        }

        return count;
    }
};
