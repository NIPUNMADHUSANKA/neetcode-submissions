class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        for ind, i in enumerate(numbers):
            x = target - i
            if x in numbers:
                return [ind+1, numbers.index(x)+1]
            