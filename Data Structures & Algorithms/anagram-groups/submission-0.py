class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        res = defaultdict(list)
        for str in strs:
            counter = defaultdict(int)
            for c in str:
                counter[c] += 1

            freq_tuple = tuple(sorted((k, v) for k, v in counter.items()))
            res[freq_tuple].append(str)
        return res.values()
            


        