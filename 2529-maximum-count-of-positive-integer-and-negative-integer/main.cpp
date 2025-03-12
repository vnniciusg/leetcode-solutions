class Solution {
public:
    int maximumCount(vector<int>& nums) {
        int neg_count = binarySearch(nums, 0);
        int pos_count = nums.size() - binarySearch(nums, 1);
        return std::max(neg_count, pos_count);
    }

private:
    int binarySearch(vector<int>& nums, int target) {
        int l = 0, r = nums.size() - 1;
        int result = nums.size();

        while (l <= r) {
            int mid = l + (r - l) / 2;

            if (nums[mid] < target) {
                l = mid + 1;
            }else {
                r = mid - 1;
                result = mid;
            }
        }

        return result;
    }
};
