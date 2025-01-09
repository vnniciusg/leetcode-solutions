class Solution {
public:
    int prefixCount(vector<string>& words, string pref) {
        int count = 0;
        
        for(string word: words) {
            if(word.starts_with(pref)) {
                count++;
            }
        }

        return count;
    }
};
