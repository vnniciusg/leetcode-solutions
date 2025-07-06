class FindSumPairs:

    def __init__(self, nums1: List[int], nums2: List[int]):
        self.nums1 = nums1
        self.nums2 = nums2    

        self.counter2 = Counter(nums2)    

    def add(self, index: int, val: int) -> None:
        org = self.nums2[index]
        self.counter2[org] -= 1
        if self.counter2[org] == 0:
            del self.counter2[org]
        
        self.nums2[index] += val
        self.counter2[self.nums2[index]] += 1

    def count(self, tot: int) -> int:
        return sum(
            self.counter2.get(tot - num, 0) for num in self.nums1
        )


# Your FindSumPairs object will be instantiated and called as such:
# obj = FindSumPairs(nums1, nums2)
# obj.add(index,val)
# param_2 = obj.count(tot)
