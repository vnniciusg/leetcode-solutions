class Solution {
public:
    bool doesValidArrayExist(vector<int>& derived) {
        int XOR = 0;

        for (const auto& number: derived) {
            XOR ^= number;
        }

        return XOR == 0;
    }
};
