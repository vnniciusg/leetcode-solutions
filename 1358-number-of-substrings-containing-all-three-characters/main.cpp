class Solution {
public:
    int numberOfSubstrings(string s) {
        int l = 0, ans = 0;
        unordered_map<char, int> count;

        int n = s.size();
        for (int r = 0; r < n; r++) {
            count[s[r]]++;

            while (count.size() == 3) {
                ans += (n - r);
                count[s[l]]--;

                if (count[s[l]] == 0) {
                    count.erase(s[l]);
                } 

                l++;
            }
        }

        return ans;
    }
};
