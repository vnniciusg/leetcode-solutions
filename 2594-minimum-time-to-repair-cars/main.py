class Solution:
    def repairCars(self, ranks: List[int], cars: int) -> int:
        
        left, right = 1, min(ranks) * (cars ** 2)

        while left < right:
            mid = left + (right - left) // 2
            cars_repaired = sum(int((mid / rank) ** 0.5) for rank in ranks)

            if cars_repaired < cars:
                left = mid +  1
            else:
                right = mid
        
        return left
