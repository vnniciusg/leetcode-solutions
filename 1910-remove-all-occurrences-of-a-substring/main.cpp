class Solution {
public:
    string removeOccurrences(string s, string part) {
        return _helper(s, part);
    }
private:
    string _helper(string substr, string part) {
        if (!substr.contains(part)) {
            return substr;
        }   
        size_t position = substr.find(part);
        if (position != std::string::npos) {
            substr.erase(position, part.length());
        }

        substr = _helper(substr, part);
        return substr;
    }
};
