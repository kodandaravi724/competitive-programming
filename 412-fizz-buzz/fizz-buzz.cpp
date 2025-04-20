class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> res(n, "");
        for(int i=0; i<n; i++) {
            string cur = "";
            if ((i+1)%3 == 0) {
                cur += "Fizz";
            }
            if ((i+1)%5 == 0) {
                cur += "Buzz";
            }
            if (cur == "") {
                res[i] = to_string((i+1));
            } else {
                res[i] = cur;
            }
        }
        return res;
    }
};