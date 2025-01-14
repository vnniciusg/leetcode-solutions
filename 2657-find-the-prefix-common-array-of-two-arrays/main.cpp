class Solution {
public:
    vector<int> findThePrefixCommonArray(vector<int>& A, vector<int>& B) {
        
        int n = A.size();
        vector<int> result(n);
        unordered_set<int> setA, setB;

        for (int i = 0; i < n; i++) {
            setA.insert(A[i]);
            setB.insert(B[i]);

            int common = 0;
            for (const auto& num: setA) {
                if (setB.count(num)) common++;
            }

            result[i] = common;
        }

        return result;
    }
};
