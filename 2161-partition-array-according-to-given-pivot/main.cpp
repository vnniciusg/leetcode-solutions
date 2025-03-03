class Solution {
public:
    vector<int> pivotArray(vector<int>& nums, int pivot) {
        vector <int> left, right, mid;

        for (const auto& num: nums) {
            if (num < pivot){
                left.push_back(num);
            }else if (num > pivot) {
                right.push_back(num);
            }else {
                mid.push_back(num);
            }
        }

        vector<int> ans;
        ans.reserve(left.size() + mid.size() + right.size());
        
        ans.insert(ans.end(), left.begin(), left.end());
        ans.insert(ans.end(), mid.begin(), mid.end());
        ans.insert(ans.end(), right.begin(), right.end());

        return ans;
    }
};
