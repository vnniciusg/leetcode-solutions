type MinHeap []int

func (h MinHeap) Len() int { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface {}) {
    *h = append(*h, x.(int))
}
func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0: n-1]
    return x
}

func minOperations(nums []int, k int) int {
    h := MinHeap(nums)
    heap.Init(&h)

    numOperations := 0

    for h[0] < k {
        
        if len(h) < 2 {
            return numOperations
        }

        smallest := heap.Pop(&h).(int)
        secondSmallest := heap.Pop(&h).(int)

        newElement := smallest*2 + secondSmallest
        heap.Push(&h, newElement)

        numOperations++
    }

    return numOperations
}
