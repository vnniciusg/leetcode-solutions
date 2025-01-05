class Solution {
public:
    string shiftingLetters(string s, vector<vector<int>>& shifts) {
        int n = s.length();
        vector<int> diff(n+1, 0);

        for (const auto& shift: shifts) {
            int start = shift[0], end = shift[1], direction = shift[2];
            if (direction == 1) {
                diff[start]++;
                diff[end + 1]--;
            } else {
                diff[start]--;
                diff[end + 1]++;
            }
        }

        int currShift = 0;
        for (int i = 0; i < n; i++) {
            currShift = ((currShift + diff[i]) % 26 + 26) % 26;
            int curr = s[i] - 'a';
            int newChar = (curr + currShift) % 26;
            s[i] = newChar + 'a';
        }

        return s;
    }
};
