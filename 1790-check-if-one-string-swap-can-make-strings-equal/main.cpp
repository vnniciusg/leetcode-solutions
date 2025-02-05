class Solution {
public:
    bool areAlmostEqual(string s1, string s2) {
        
        if (s1 == s2) return true;
        
        unordered_map<char, int> freq1, freq2;
        for (char c: s1) freq1[c]++;
        for (char c: s2) freq2[c]++;
        
        if (freq1 != freq2) return false;

        int diff = 0;
        for (size_t i = 0; i < s1.size(); i++) {
            if (s1[i] != s2[i]) diff++;
        }
    
        return diff == 2;
    }
};
