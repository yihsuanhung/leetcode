class Solution:
    def largestAltitude(self, gain: List[int]) -> int:
        current_altitude, max_altitude = 0, 0
        for a in gain:
            current_altitude += a
            max_altitude = max(max_altitude, current_altitude)
        return max_altitude
