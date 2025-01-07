class Solution {
public:
    vector<string> stringMatching(vector<string>& words) {
        string str;
        for (const string& word: words) {
            str += word + " ";
        }

        vector<string> ans;
        unordered_set<string> added;

        for (const string& word: words) {
            size_t count = 0;
            size_t pos = str.find(word, 0);
            while (pos != string::npos) {
                count++;
                pos = str.find(word, pos + 1);
            }

            if (count > 1 && added.find(word) == added.end()) {
                ans.push_back(word);
                added.insert(word);
            }
        }

        return ans;
    }
};
