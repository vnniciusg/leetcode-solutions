class Solution:
    def canPlaceFlowers(self, flowerbed: List[int], n: int) -> bool:
        
        extended_flowerbed = [0] + flowerbed + [0]
        planted_flowers = 0

        for idx in range(1, len(extended_flowerbed) - 1):

            if extended_flowerbed[idx-1] == 0 and extended_flowerbed[idx] == 0 and extended_flowerbed[idx + 1] == 0:
                extended_flowerbed[idx] = 1
                planted_flowers+= 1

        return planted_flowers >= n
