func sortColors(nums []int)  {
    
    low, mid, high := 0, 0, len(nums) - 1

    for mid <= high {
        if nums[mid] == 0 {
            nums[low], nums[mid] = nums[mid], nums[low]
            low++
            mid++
        }else if nums[mid] == 1 {
            mid++
        }else {
            nums[high], nums[mid] = nums[mid], nums[high]
            high--
        }
    }
}
