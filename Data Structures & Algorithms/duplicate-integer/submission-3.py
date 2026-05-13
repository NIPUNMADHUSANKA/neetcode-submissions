class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        list_hash = set()

        for i in nums:
            if i in list_hash:
                return True
            list_hash.add(i)
        return False