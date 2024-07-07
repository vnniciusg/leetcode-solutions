func numWaterBottles(numBottles int, numExchange int) int {

    totalDrunk, emptyBottles := numBottles, numBottles

    for emptyBottles >= numExchange {

        newFullBottles := emptyBottles / numExchange
        totalDrunk += newFullBottles

        emptyBottles = emptyBottles % numExchange + newFullBottles
    }   

    return totalDrunk
}
