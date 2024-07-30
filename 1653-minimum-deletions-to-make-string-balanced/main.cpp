class Solution {
public:
    int minimumDeletions(string s) {

        int b_count = 0, deletions = 0;
        
        for (char chr: s) {
            if (chr == 'b'){ b_count++; }
            else {
                deletions = std::min(deletions + 1, b_count);
            }
        }

        return deletions;
    }
};
