from typing import List
class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        res = []
        maxcand = max(candies)
        for i in candies:
            if i + extraCandies >= maxcand:
                res.append(True)
            else:
                res.append(False)
        return res
    
solution = Solution()

# Тестовые случаи
test_cases = [
    ([2, 3, 5, 1, 3], 3),  # Ожидаемый результат: [True, True, True, False, True]
    ([4, 2, 1, 1, 2], 1),  # Ожидаемый результат: [True, False, False, False, False]
    ([12, 1, 12], 10),     # Ожидаемый результат: [True, False, True]
]

# Проверяем каждый тестовый случай
for candies, extra in test_cases:
    result = solution.kidsWithCandies(candies, extra)
    print(f'candies: {candies}, extraCandies: {extra} => result: {result}')