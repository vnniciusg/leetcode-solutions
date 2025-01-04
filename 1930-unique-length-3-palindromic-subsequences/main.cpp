class Solution {
public:
    int countPalindromicSubsequence(string s) {
        unordered_set<char> letters(s.begin(), s.end());
        int ans = 0;

        for(char letter: letters) {
            int i = s.find(letter);
            int j = s.rfind(letter);

            if (i == j || i == string::npos) continue;

            unordered_set<char> between;
            for (int k = i + 1; k < j; k++) {
                between.insert(s[k]);
            }

            ans += between.size();
        }
        
        return ans;
    }
};
