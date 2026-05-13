class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hash_arr = {}

        for index, i in enumerate(nums):
            compute = target - i
            if compute in hash_arr:
                return [hash_arr[compute], index]
            hash_arr[i] = index