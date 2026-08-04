class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        if len(s) != len(t):
            return False
        
        count_s = {}
        count_t = {}
        
        # Count characters in s
        for ch in s:
            count_s[ch] = count_s.get(ch, 0) + 1
        
        # Count characters in t
        for ch in t:
            count_t[ch] = count_t.get(ch, 0) + 1
        
        return count_s == count_t
