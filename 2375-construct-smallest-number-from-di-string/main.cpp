class Solution {
public:
    string smallestNumber(string pattern) {
        size_t n = pattern.size();
        vector<string> ans;
        vector<int> numStack;

        for (size_t idx = 0; idx < n + 1; idx++) {

            numStack.push_back(idx + 1);

            if (idx == n || pattern[idx] == 'I') {
                while ( numStack.size() > 0 ) {
                    int top = numStack.back();
                    numStack.pop_back();
                    ans.push_back(to_string(top));
                }
            }
        }

        return accumulate(ans.begin(), ans.end(), string());
    }
};
