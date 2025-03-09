class Solution {
public:
    int numberOfAlternatingGroups(vector<int>& colors, int k) {
        
        int n = colors.size();
        int ans = 0, l = 0;

        for (int r = 1; r < n + k - 1; r++) {
            if (colors[r % n] == colors[(r - 1) % n]) {
                l = r;
            }

            if (r - l + 1 > k) {
                l++;
            } 

            if (r - l + 1 == k) {
                ans++;
            }
        }

        return ans;
    }
};
