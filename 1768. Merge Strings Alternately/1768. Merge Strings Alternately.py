class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        merged = ""
        len1 = len(word1)
        len2 = len(word2)
        if len1 > len2:
            gotofirstword = True
        else:
            gotofirstword = False
        if gotofirstword == True:
            for i in range(len1):
                merged += word1[i]
                try:
                    merged += word2[i]
                except:
                    pass
        else:
            for i in range(len2):
                try:
                    merged += word1[i]
                except:
                    pass
                merged += word2[i]
                
        return merged

solution = Solution()

# Тестируем функцию
test_cases = [
    ("ab", "pqrs"),  # Ожидаемый результат: "apbqcr"
    ("ab", "pqrs"),  # Ожидаемый результат: "apbqrs"
    ("abcd", "pq"),  # Ожидаемый результат: "apbqcd"
    ("", "xyz"),     # Ожидаемый результат: "xyz"
    ("abc", ""),     # Ожидаемый результат: "abc"
    ("", "")         # Ожидаемый результат: ""
]


for word1, word2 in test_cases:
    result = solution.mergeAlternately(word1, word2)
    print(f"mergeAlternately('{word1}', '{word2}') = '{result}'")