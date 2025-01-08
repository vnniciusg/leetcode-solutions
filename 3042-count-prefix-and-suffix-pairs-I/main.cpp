class Solution {
public:
    int countPrefixSuffixPairs(vector<string>& words) {
        size_t n = words.size();
        int count = 0;

        for (size_t i = 0; i < n; i++) {
            for (size_t j = i + 1; j < n; j++) {
                if (words[j].starts_with(words[i]) && words[j].ends_with(words[i])) {
                    count++;
                }
            }
        }

        return count;
        
    }
};
