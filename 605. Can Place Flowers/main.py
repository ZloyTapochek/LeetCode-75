from typing import List

class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        count = 0
        lenght = len(flowerbed)
        for i in range(lenght):
            if flowerbed[i] == 0:
                if (i == 0 or flowerbed[i - 1] == 0) and (i == lenght - 1 or flowerbed[i + 1] == 0):
                    flowerbed[i] = 1
                    count+=1
        return count >= n

def test_canPlaceFlowers():
    solution = Solution()
    
    # Тест 1: Простой случай
    assert solution.canPlaceFlowers([1, 0, 0, 0, 1], 1) == True, "Test case 1 failed"
    
    # Тест 2: Невозможно посадить цветы
    assert solution.canPlaceFlowers([1, 0, 0, 0, 1], 2) == False, "Test case 2 failed"
    
    # Тест 3: Можно посадить цветы
    assert solution.canPlaceFlowers([0, 0, 1, 0, 0], 1) == True, "Test case 3 failed"
    
    # Тест 4: Можно посадить несколько цветов
    assert solution.canPlaceFlowers([0, 0, 0, 0, 0], 3) == True, "Test case 4 failed"
    
    # Тест 5: Невозможно посадить нужное количество цветов
    assert solution.canPlaceFlowers([0, 0, 0, 0, 0], 4) == False, "Test case 5 failed"
    
    # Тест 6: Пустой цветник
    assert solution.canPlaceFlowers([], 1) == False, "Test case 6 failed"
    
    print("All test cases passed!")

# Запуск тестов
test_canPlaceFlowers()

                