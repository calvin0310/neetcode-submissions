class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        s_count = {}
        t_count = {}
        for i in s:
            if s_count.get(i, 0) == 0:
                s_count[i] = 1
            else:
                s_count[i] += 1
        for j in t:
            if t_count.get(j, 0) == 0:
                t_count[j] = 1
            else:
                t_count[j] += 1
        if s_count == t_count:
            return True
        else:
            return False