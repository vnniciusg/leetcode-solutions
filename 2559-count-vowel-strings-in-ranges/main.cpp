class Solution {
public:
    vector<int> vowelStrings(vector<string>& words, vector<vector<int>>& queries) {
        int n = words.size();
        vector<int> prefix_sum(n);
        int _sum = 0;

        for (int i = 0; i < n; i++) {
            string curr_word = words[i];
            if (isVowel(curr_word[0]) && isVowel(curr_word.back())) {
                _sum++;
            }
            prefix_sum[i] = _sum;
        }

        n = queries.size();
        vector<int> ans(n);
        for (int i = 0; i < n; i++ ){
            vector<int> curr_query = queries[i];
            ans[i] = prefix_sum[curr_query[1]] - (curr_query[0] == 0 ? 0 : prefix_sum[curr_query[0] - 1]);
        }

        return ans;
    }

private:
    bool isVowel(char letter) {
        static const unordered_set<char> vowels{'a', 'e', 'i', 'o', 'u'};
        return vowels.count(tolower(letter)) > 0;    
    }
};
