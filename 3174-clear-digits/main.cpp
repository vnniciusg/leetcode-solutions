class Solution {
public:
    string clearDigits(string s) {
        vector<char> ans;

        for (char ch : s) {
            if (isdigit(ch) && !ans.empty()) {
                ans.pop_back();
            } else {
                ans.push_back(ch);
            }
        }

        return string(ans.begin(), ans.end());
    }
};
