class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        seen = {}
        for ind, i in enumerate(nums):
            com = target - i
            if com in seen:
                return [seen[com], ind]
            seen[i] = ind