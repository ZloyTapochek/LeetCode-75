class Solution:
    def pivotIndex(self, nums: List[int]) -> int:
        # число индекса
        Allsum = sum(nums)
        Leftsum = 0
        Rigthsum = Allsum - Leftsum

        if Leftsum == Rigthsum - nums[0]:
            return 0
        
        for i in range(len(nums)):
            curr_int = nums[i]
            Rigthsum = Rigthsum - curr_int
            if Leftsum == Rigthsum:
                return i
            Leftsum = Leftsum+curr_int
        
        return -1