class Solution {
public:
    vector<int> findMissingAndRepeatedValues(vector<vector<int>>& grid) {
        int n = grid.size();
        unordered_map<int, int> count;

        for (const auto& row: grid) {
            for (int num: row) {
                count[num]++;
            }
        }

        int repeated = -1, missing = -1;
        for (int num = 1; num <= n * n; num++) {
            if (count[num] == 2) {
                repeated = num;
            }else if (count[num] == 0) {
                missing = num;
            }
        }

        return {repeated, missing};
    }
};
