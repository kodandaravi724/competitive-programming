class Solution {
public:
    int longestPalindrome(string s) {
        bool isOdd = false;
        unordered_map<char, int> freq;
        for(char& c: s) {
            freq[c]++;
        }
        int longestPalindrome = 0;
        for(auto &[k, v]: freq) {
            if (v&1 == 1) {
                if (!isOdd) {
                    isOdd = true;
                }
                longestPalindrome += (v-1);
            } else {
                longestPalindrome += v;
            }
        }
        return longestPalindrome + isOdd;
    }
};