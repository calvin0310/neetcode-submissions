class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        duplicated_map = {}
        for num in nums:
            # num = str(num)
            if duplicated_map.get(num, 0) == 0:
                duplicated_map[num] = 1
            else:
                duplicated_map[num] += 1
            if duplicated_map[num] > 1:
                return True
        return False