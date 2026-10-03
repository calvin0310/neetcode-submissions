class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        differenceMap = {}
        for i in range(len(nums)):
            difference = target - nums[i]
            print(difference, differenceMap)

            if difference in differenceMap:
                return [differenceMap[difference], i]
            else:
                differenceMap[nums[i]] = i

