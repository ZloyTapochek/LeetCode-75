class Solution:
    def longestSubarray(self, nums: List[int]) -> int:
        left, delets, maxl = 0,0,0
        for right in range(len(nums)):
            if nums[right] == 0:
                delets += 1
            while delets > 1:
                if nums[left] == 0:
                    delets -=1
                left += 1
            maxl = max(maxl, right-left)
        return maxl
    
# Важно:
# Т.к. мы ОБЯЗАНЫ удалить элемент, а только ПОСЛЕ этого вернуть длину массива.
# То мы будем возвращать длину массива `right-left` без включения левого элемента.