func canPlaceFlowers(flowerbed []int, n int) bool {
    
    extendedFlowerbed := append([]int{0}, flowerbed...)
    extendedFlowerbed = append(extendedFlowerbed, 0)
    plantedFlowers := 0

    for idx := 1; idx < len(extendedFlowerbed) - 1; idx++ {
        if extendedFlowerbed[idx - 1] == 0 && extendedFlowerbed[idx] == 0 && extendedFlowerbed[idx + 1] == 0 {
            extendedFlowerbed[idx] = 1
            plantedFlowers++
        }
    }

    return plantedFlowers >= n
}
