class Solution {
public:
    int minimumLength(string s) {
        unordered_map<char, int> freq_map;
        for (const char& chr: s) 
            freq_map[chr]++;

        int delete_count = 0;
        for (const auto& pair: freq_map) {
            int frequency = pair.second;
            if (frequency & 1 != 0) {
                delete_count += frequency - 1;
            }else {
                delete_count += frequency - 2;
            }
        }

        return s.size() - delete_count;
    }
};
