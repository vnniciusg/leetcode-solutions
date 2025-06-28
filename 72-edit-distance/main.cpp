class Solution {
public:
    int minDistance(string word1, string word2) {
        
        size_t m = word1.size();
        size_t n = word2.size();

        vector<vector<int>> dp(m + 1, vector<int>(n + 1));
        for (size_t i = 0; i <= m; i++){
            dp[i][0] = i;
        }
        for (size_t j = 0; j <= n; j++) {
            dp[0][j] = j;
        }

        for (size_t i = 1; i < m + 1; i++) {
            for (size_t j = 1; j < n + 1; j++){
                int cost = word1[i - 1] == word2[j - 1] ? 0 : 1;
                dp[i][j] = min({
                    dp[i - 1][j] + 1,
                    dp[i][j - 1] + 1,
                    dp[i - 1][j - 1] + cost,
                });
            }
        }

        return dp[m][n];
    }
};
