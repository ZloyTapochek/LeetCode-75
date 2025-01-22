
class Solution:
    def longestOnes(self, nums: list[int], k: int) -> int:
        left, maxl, zeroes = 0, 0, 0
        for right in range(len(nums)):
            if nums[right] == 0:
                zeroes += 1
            while zeroes > k:
                if nums[left] == 0:
                    zeroes -= 1
                left += 1
            maxl = max(maxl, len(nums[left:right+1]))
        return maxl

solution = Solution()

nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1]
k = 3
print(solution.longestOnes(nums, k))