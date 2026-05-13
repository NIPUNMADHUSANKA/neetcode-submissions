class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        
        str1 = {}
        str2 = {}

        for i in s:
            if str1.get(i):
                str1[i] = str1[i] + 1
            else:
                str1[i] = 1
        
        for i in t:
            if str2.get(i):
                str2[i] = str2[i] + 1
            else:
                str2[i] = 1
        
        if str1 == str2:
            return True
        else:
            return False