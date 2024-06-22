class Solution {
public:
    int kthFactor(int n, int k) {
        
        vector<int> factor_list;
        for (int i = 1; i < n + 1; i++) {
            if ((n % i) == 0) {
                factor_list.push_back(i);
            }
        }

        return (k - 1) >= factor_list.size() ? -1 : factor_list[k - 1];
    }
};
