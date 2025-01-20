class Solution:
    def maxVowels(self, s: str, k: int) -> int:
        vowels = ['a', 'e', 'i', 'o', 'u']
        length = len(s)
        if length <= k:
            count = 0
            for i in s:
                if i in vowels:
                    count+=1
            return count
        else:
            max_count = 0
            temp_count = 0
            j = 0
            while j+k<=length:
                for i in range(j, j+k):
                    if s[i] in vowels:
                        temp_count +=1
                max_count = max(temp_count, max_count)
                temp_count = 0
                j+=1
            return max_count
        

soliton = Solution()

s = "weallloveyou"
k = 7
print(soliton.maxVowels(s,k))