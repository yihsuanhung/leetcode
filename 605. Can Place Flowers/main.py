class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        count = 0
        for i in range(len(flowerbed)):
            if flowerbed[i] == 0:
                # check if both left and right sides are empty
                is_left_empty = True if i == 0 else flowerbed[i-1] == 0
                is_right_empty = True if i == len(
                    flowerbed)-1 else flowerbed[i+1] == 0
                if is_left_empty and is_right_empty:
                    flowerbed[i] = 1
                    count += 1
        return count >= n
