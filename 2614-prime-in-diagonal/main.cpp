oclass Solution {
public:
    int diagonalPrime(vector<vector<int>>& nums) {
        
        if (nums.empty()) {
            return 0;
        }

        vector<int> prime_numbers;
        prime_numbers.reserve(nums.size() * 2);

        for (size_t i = 0; i < nums.size(); ++i) {
            prime_numbers.push_back(nums[i][i]);
            prime_numbers.push_back(nums[i][nums.size() - i - 1]);
        }

        int max_value = 0;
        for (int num: prime_numbers) {
            if (is_prime(num) && num > max_value){
                max_value = num;
            }
        }

        return max_value;

    }

private:
    bool is_prime(int num){
        if (num <= 1) {
            return false;
        }

        for (int i = 2; i <= sqrt(num); ++i) {
            if (num % i == 0){
                return false;
            }
        }

        return true;
    }
};
