class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        res1,res2 = [],[]
        for i in range(len(nums1)):
            if (nums1[i] in nums2) or (nums1[i] in res2):
                continue
            else:
                res2.append(nums1[i])
        for i in range(len(nums2)):
            if (nums2[i] in nums1) or (nums2[i] in res1):
                continue
            else:
                res1.append(nums2[i])
        return res2,res1
        



# Решение:
class Solution:
    def findDifference(self, nums1: List[int], nums2: List[int]) -> List[List[int]]:
        l = []
        l.append(list(set(nums1) - set(nums2)))
        l.append(list(set(nums2) - set(nums1)))
        return l
