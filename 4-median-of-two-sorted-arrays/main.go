package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

    var arr []int
    arr = append(arr, nums1...)
    arr = append(arr, nums2...)

    sort.Slice(arr, func(i, j  int) bool {
        return arr[i] < arr[j]
    })

    mid := len(arr) / 2
    if len(arr) & 1 != 0 {
        return float64(arr[mid])
    }

    return float64(arr[mid] + arr[mid - 1]) / 2

}
