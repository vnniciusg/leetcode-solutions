type FindSumPairs struct {
    nums1 []int
    nums2 []int

    counter2 map[int]int
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {

    counter := make(map[int]int)
	for _, num := range nums2 {
		counter[num]++
	}

	return FindSumPairs{
		nums1:    nums1,
		nums2:    nums2,
		counter2: counter,
	}
}

func (this *FindSumPairs) Add(index int, val int)  {
    original := this.nums2[index]
	this.counter2[original]--
	if this.counter2[original] == 0 {
		delete(this.counter2, original)
	}

	this.nums2[index] += val
	this.counter2[this.nums2[index]]++
}

func (this *FindSumPairs) Count(tot int) int {
    count := 0
	for _, num := range this.nums1 {
		need := tot - num
		if c, exists := this.counter2[need]; exists {
			count += c
		}
	}
	return count
}


/**
 * Your FindSumPairs object will be instantiated and called as such:
 * obj := Constructor(nums1, nums2);
 * obj.Add(index,val);
 * param_2 := obj.Count(tot);
 */
