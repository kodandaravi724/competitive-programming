class Solution {
public:
    int dfs(vector<int>& nums, const int& target, int& curSum, int& combinations, vector<int>& memo) {
        if (curSum == target) {
            return 1;
        }
        if (curSum > target) {
            return 0;
        }
        if ( memo[target-curSum] != -1 ) {
            return memo[target-curSum];
        }
        int count = 0;
        for(int i = 0; i < nums.size(); i++) {
            curSum += nums[i];
            count += dfs(nums, target, curSum, combinations, memo);
            curSum -= nums[i];
        }
        memo[target-curSum] = count;
        return memo[target-curSum];
    }
    int combinationSum4(vector<int>& nums, int target) {
        int combinations = 0;
        int curSum = 0;
        vector<int> memo(target+1, -1);
        return dfs(nums, target, curSum, combinations, memo);
    }
};