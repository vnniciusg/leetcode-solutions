class Solution:
    def checkRecord(self, s: str) -> bool:
        absences = s.count('A')
        if absences >= 2:
            return False
        
        if 'LLL' in s:
            return False
        
        return True
