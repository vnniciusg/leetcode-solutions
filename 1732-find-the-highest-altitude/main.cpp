class Solution {
public:
    int largestAltitude(vector<int>& gain) {
        
        int max_altitude = 0;
        int curr_altitude = 0;
        
        for (const auto& altitude: gain) {
            curr_altitude += altitude;
            max_altitude = max(max_altitude, curr_altitude);
        }
        
        return max_altitude;
    }
};
