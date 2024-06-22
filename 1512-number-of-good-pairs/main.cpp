class Solution {
public:
    int numIdenticalPairs(vector<int>& nums) {
        
        std::unordered_map<int, int> count_values;
        int count_good_pairs = 0;

        for (const int& num : nums) {
            if (count_values.find(num) != count_values.end()) {
                count_good_pairs += count_values[num];
            }
            count_values[num]++;
        }
        
        return count_good_pairs;
    }
};
