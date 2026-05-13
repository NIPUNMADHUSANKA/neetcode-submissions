class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        list2 = set()

        for i in nums:
            if i in list2:
                return True
            else:
                list2.add(i)
        return False