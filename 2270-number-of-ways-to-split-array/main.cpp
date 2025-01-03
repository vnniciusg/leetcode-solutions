class Solution {
public:
    int waysToSplitArray(vector<int>& nums) {
        long left = 0;
        long right = sum(nums);

        int count = 0;
        for (int i = 0; i < nums.size() - 1; i++) {
            left += nums[i];
            right -= nums[i];

            if (left >= right) {
                count++;
            }
        }

        return count;
    }

private:
    long sum(vector<int>& nums) {
        long _sum = 0;
        for(auto num: nums) {
            _sum += num;
        }
        return _sum;
    }
};
