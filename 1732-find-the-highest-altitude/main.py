class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        
        max_altitude = curr_altitude = 0

        for altitude in gain:
            curr_altitude += altitude
            max_altitude = max(max_altitude, curr_altitude)
        

        return max_altitude
