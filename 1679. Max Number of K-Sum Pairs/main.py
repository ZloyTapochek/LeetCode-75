class Solution:
    def maxOperations(self, nums: list[int], k: int) -> int:
        count_of_operations = 0
        nums.sort()
        i,j = 0, len(nums)-1
        while(i<j):
            if nums[i]+nums[j] == k:
                nums.pop(j)
                
                nums.pop(i)
                j-=2
                count_of_operations+=1
            else:
                if nums[i]+nums[j] >k:
                    j-=1
                else:
                    i+=1
        return count_of_operations
    
solution = Solution()

nums = [4,4,1,3,1,3,2,2,5,5,1,5,2,1,2,3,5,4]
k = 2
result = solution.maxOperations(nums, k)

# Выводим результат
print(result)  # Вывод: 2 (пары: (1, 4) и (2, 3))